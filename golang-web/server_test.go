package main

import (
	"fmt"
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	server := http.Server{
		Addr: "localhost:3000",
	}
	
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestServeMux(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, world!")
	})
	mux.HandleFunc("/goodbye", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Goodbye, world!")
	})

	server := http.Server{
		Addr: "localhost:3000",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestServeMux2(t *testing.T) {
	mux := http.NewServeMux()
	mux.Handle("/hello", &helloHandler{})
	mux.Handle("/goodbye", &goodbyeHandler{})

	server := http.Server{
		Addr: "localhost:3000",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

type helloHandler struct{}
func (h *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

type goodbyeHandler struct{}
func (h *goodbyeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Goodbye, world!")
}	

