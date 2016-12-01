package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
)

type Response struct {
	Code   int    `json:"code"`
	Result string `json:"result"`
}

// Docs
// https://golang.org/pkg/net/http
// https://golang.org/pkg/io/#Writer

// This is our function we are going to use to handle the request
// All handlers need to accept two arguments
// 1. Is the ResponseWriter interface, we use this to write a reponse back to the client
// 2. Is the Reponse struct which holds useful information about the request headers, method, url etc
func hello(w http.ResponseWriter, r *http.Request) {

	name := r.FormValue("name")
	httpCode := 200
	result := fmt.Sprintf("%s %s", "hello", name)

	switch {
	case len(name) == 1:
		httpCode = 400
		result = "Name must greater the one character long"
	case name == "":
		httpCode = 400
		result = "Provide your name as argument e.g ?name=<name>"
	}

	sendResponse(httpCode, result, w)
}

func sendResponse(code int, respond string, w http.ResponseWriter) {
	res := Response{
		Code:   code,
		Result: respond,
	}
	json, err := json.Marshal(res)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write(json)
}

func main() {
	// Add ads the function thats going to handle that response
	http.HandleFunc("/", hello)
	// Starts the web server
	// The first argument in this method is the port you want your server to run on
	// The second is a handler. However we have already added this in the line above so we just pass in nil

	// Specify the default Port
	port := flag.String("port", "8000", "an string")
	flag.Parse()

	fmt.Println("Server Started in port :", *port)
	fmt.Println("To change the port use --port=<number>")

	http.ListenAndServe(":"+*port, nil)
}
