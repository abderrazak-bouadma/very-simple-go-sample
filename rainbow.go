package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func handleRootContext(w http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(w, "<!doctype html>"+
		"<html>"+
		"<head>"+
		"<meta charset=\"utf-8\">"+
		"<title>A static page</title>"+
		"<style>"+
		"body {background-color: yellowgreen}"+
		"</style>"+
		"</head>"+
		"<body>"+
		"<h1>Hello from a static page</h1>"+
		"</body>"+
		"</html>")
}

func handleHealthCheck(w http.ResponseWriter, request *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"alive": true}`)
}

func main() {
	http.HandleFunc("/", handleRootContext)
	http.HandleFunc("/health", handleHealthCheck)
	log.Println("Listening ....")
	http.ListenAndServe(":3000", nil)
}
