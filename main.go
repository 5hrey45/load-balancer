package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

var (
	servers = []string{
		"http://127.0.0.1:5000",
		"http://127.0.0.1:5001",
		"http://127.0.0.1:5002",
		"http://127.0.0.1:5003",
		"http://127.0.0.1:5004",
	}
	lastServedIndex = 0
)

func main() {
	http.HandleFunc("/", forwardRequest)
	http.ListenAndServe(":8080", nil)
}

func forwardRequest(rw http.ResponseWriter, req *http.Request) {
	url := getServer()
	reverseProxy := httputil.NewSingleHostReverseProxy(url)
	fmt.Println("Routing request to the server", url)
	reverseProxy.ServeHTTP(rw, req)
}

func getServer() *url.URL {
	nextIndex := (lastServedIndex + 1) % 5
	url, err := url.Parse(servers[lastServedIndex])
	lastServedIndex = nextIndex
	if err != nil {
		log.Fatal(err)
	}
	return url
}