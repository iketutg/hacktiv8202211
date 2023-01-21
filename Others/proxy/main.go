package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

const (
	PORT = "8080"
)

func logRequest(proxyURL string) {
	log.Printf("proxy_url: %s\n", proxyURL)
}
func serveProxy(targeturl string, res http.ResponseWriter, req *http.Request) {

	url, _ := url.Parse(targeturl)

	proxy := httputil.NewSingleHostReverseProxy(url)

	proxy.ServeHTTP(res, req)
}
func handlerServer1(res http.ResponseWriter, req *http.Request) {
	log.Println("Server1 Forward")
	url := "http://localhost:8181/server1"
	serveProxy(url, res, req)
}

func handlerServer2(res http.ResponseWriter, req *http.Request) {
	log.Println("Server1 Forward")
	url := "http://localhost:8181/server1"
	serveProxy(url, res, req)
}

func main() {
	// start server
	http.HandleFunc("/server1", handlerServer1)
	http.HandleFunc("/server2", handlerServer2)

	log.Fatal(http.ListenAndServe(":"+PORT, nil))
}
