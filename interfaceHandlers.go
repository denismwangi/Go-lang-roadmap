package main

import (
	"fmt"
	"net/http"
)

// Handler is an interface with a method signature that matches http.Handler's ServeHTTP method
type Handler interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
}

// MyHandler is a custom struct that implements the Handler interface
type MyHandler struct{}

// ServeHTTP is the method that satisfies the http.Handler interface
func (h MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world")
}

func main() {
	// Using the custom handler
	handler := MyHandler{}
	http.ListenAndServe(":9091", handler)
}
