package main

import (
	"bytes"
	"io"

	"openceptor.eu/config"
	"openceptor.eu/connection"
	"openceptor.eu/handler"
	"openceptor.eu/project"
	"openceptor.eu/request"

	"github.com/ilyakaznacheev/cleanenv"

	"errors"
	"flag"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

// Args command-line parameters
type Args struct {
	ConfigPath string
}

func generateServerError(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Openceptor", "Invalid Application ID")
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(`{"error": "invalid application ID for Openceptor API"}`))
}

func serve(w http.ResponseWriter, r *http.Request) {
	re := regexp.MustCompile(`(?i)^\/([0-9A-F]{8}-[0-9A-F]{4}-4[0-9A-F]{3}-[89AB][0-9A-F]{3}-[0-9A-F]{12})`)
	path := r.URL.Path

	id := re.FindSubmatch([]byte(path))
	if nil == id {
		generateServerError(w)
		return
	}

	var project project.Project
	project.Load(string(id[1]))

	// Read the body in first to be able to manipulate twice the content of request
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error reading body: %s", err), 500)
		return
	}

	// Send Request into RabbitMQ for the web application
	r.Body = io.NopCloser(bytes.NewBuffer(body))
	requestObject, _ := request.CreateRequestFromHttpRequest(r, &cfg, project.Id)
	go handler.SendToQueue(&requestObject)

	// Get Mocking Rule if it exists
	mockingRule := project.GetMockingRule(r.Method, r.RequestURI)
	if mockingRule != nil {
		// Request is mocking, serve the mocking response
		for headerIndex, headerValue := range mockingRule.ResponseHeaders {
			w.Header().Set(headerIndex, headerValue.(string))
		}
		w.Write([]byte(mockingRule.ResponseBody))
		if mockingRule.ResponseStatus != 200 {
			w.WriteHeader(int(mockingRule.ResponseStatus))
		}

		return
	}

	// Proxy the request if endpoint is defined
	if project.Endpoint != "" {
		r.Body = io.NopCloser(bytes.NewBuffer(body))
		handler.HandleRequestAndRedirect(w, r, project.Endpoint)
	}

	fmt.Fprintln(w, "Hello, world."+string(id[1]))
	handler.Test()
}

var cfg config.Config

func main() {
	args := ProcessArgs(&cfg)

	// read configuration from the file and environment variables
	if _, err := os.Stat(args.ConfigPath); errors.Is(err, os.ErrNotExist) {
		if err := cleanenv.ReadEnv(&cfg); err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
	} else {
		if err := cleanenv.ReadConfig(args.ConfigPath, &cfg); err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
	}

	// Initialize database
	db := connection.GetDbInstance(&cfg)
	defer db.Close()

	// Initialize RabbitMQ
	conn := connection.GetRabbitMqInstance(&cfg)
	defer conn.Close()

	// Serve HTTP request
	http.HandleFunc("/", serve)

	fmt.Println("Start listen to " + cfg.Server.Host + ":" + strconv.Itoa(cfg.Server.Port))
	http.ListenAndServe(cfg.Server.Host+":"+strconv.Itoa(cfg.Server.Port), nil)
}

func ProcessArgs(cfg interface{}) Args {
	var a Args

	f := flag.NewFlagSet("Openceptor", 1)

	f.StringVar(&a.ConfigPath, "c", "config.yml", "Path to configuration file")

	fu := f.Usage
	f.Usage = func() {
		fu()
		envHelp, _ := cleanenv.GetDescription(cfg, nil)
		fmt.Fprintln(f.Output())
		fmt.Fprintln(f.Output(), envHelp)
	}

	f.Parse(os.Args[1:])

	return a
}
