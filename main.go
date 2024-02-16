package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	http.HandleFunc("/", forwardRequest)
	http.ListenAndServe(":8080", nil)
}

func forwardRequest(rw http.ResponseWriter, req *http.Request) {
	url := getServer()
	reverseProxy := httputil.NewSingleHostReverseProxy(url)
	reverseProxy.ServeHTTP(rw, req)
}

func getServer() *url.URL {
	url, err := url.Parse("http://127.0.0.1:5000")
	if err != nil {
		log.Fatal(err)
	}
	return url
}