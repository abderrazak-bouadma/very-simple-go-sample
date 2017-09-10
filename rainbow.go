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

<<<<<<< HEAD
	log.Println("Listening on port 3000...")
=======
func handleHealthCheck(w http.ResponseWriter, request *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"alive": true}`)
}

func main() {
	http.HandleFunc("/", handleRootContext)
	http.HandleFunc("/health", handleHealthCheck)
	log.Println("Listening ....")
>>>>>>> 99db6eea9e16dbf54480ffb69b0ab402532db064
	http.ListenAndServe(":3000", nil)
}
