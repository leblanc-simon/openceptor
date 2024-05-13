package handler

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func HandleRequestAndRedirect(res http.ResponseWriter, r *http.Request, endpoint string) {
	url, err := url.Parse(endpoint)
	if err != nil {
		fmt.Println("Errored when sending request to the server")
		return
	}

	proxy := httputil.NewSingleHostReverseProxy(url)
	proxy.ServeHTTP(res, r)
}
