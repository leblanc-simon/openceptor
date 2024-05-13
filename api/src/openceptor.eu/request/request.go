package request

import (
	"openceptor.eu/config"

	"encoding/base64"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Request struct {
	ProjectId   string                   `json:"projectId"`
	HttpVersion string                   `json:"httpVersion"`
	Host        string                   `json:"host"`
	Method      string                   `json:"method"`
	RequestURI  string                   `json:"requestUri"`
	Headers     map[string][]string      `json:"headers"`
	Cookies     []string                 `json:"cookies"`
	Querystring string                   `json:"querystring"`
	PostValues  string                   `json:"postValues"`
	FormValues  map[string][]string      `json:"formValues"`
	FormFiles   map[string][]RequestFile `json:"formFiles"`
	RemoteAddr  string                   `json:"remoteAddr"`
}

type RequestFile struct {
	Filename    string              `json:"filename"`
	Header      map[string][]string `json:"header"`
	Size        int64               `json:"size"`
	FileContent string              `json:"content"`
}

func CreateRequestFromHttpRequest(r *http.Request, c *config.Config, projectId string) (Request, error) {
	var request Request

	request.ProjectId = projectId
	request.HttpVersion = r.Proto
	request.Host = r.Host
	request.Method = r.Method
	request.Headers = r.Header
	request.RequestURI = r.URL.Path
	request.Querystring = r.URL.RawQuery
	request.RemoteAddr = r.RemoteAddr

	for _, cookie := range r.Cookies() {
		request.Cookies = append(request.Cookies, cookie.String())
	}

	if r.Method == "POST" {
		r.ParseForm()
		request.PostValues = r.PostForm.Encode()
	}

	if r.Method == "POST" {
		err := r.ParseMultipartForm(int64(c.Server.UploadMaxSize))

		if err != nil {
			log.Fatal("ooopsss an error occurred, please try again")
		}

		// @todo check multipartform before use value
		if len(r.MultipartForm.Value) > 0 {
			request.FormValues = r.MultipartForm.Value
		}

		if len(r.MultipartForm.File) > 0 {
			for key, files := range r.MultipartForm.File {
				request.FormFiles = make(map[string][]RequestFile)
				for _, file := range files {
					fileHandler, _ := file.Open()

					data, _ := io.ReadAll(fileHandler)
					fileContent := base64.StdEncoding.EncodeToString(data)
					request.FormFiles[key] = append(request.FormFiles[key], RequestFile{
						Filename:    file.Filename,
						Header:      file.Header,
						Size:        file.Size,
						FileContent: fileContent,
					})
				}
			}
		}
	}

	return request, nil
}

func (r *Request) ToJson() ([]byte, error) {
	return json.Marshal(r)
}
