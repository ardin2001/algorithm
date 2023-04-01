package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	var message string
	mux.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		message = "root web"
		w.Write([]byte(message))
	})
	mux.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		name := r.URL.Query().Get("name")
		if r.Method == "GET" {
			message = "get products " + id + name
			w.Write([]byte(message))
		} else if r.Method == "DELETE" {
			message = "delete products " + id + name
			w.Write([]byte(message))
		}
	})
	http.HandleFunc("/test", func(w http.ResponseWriter, _ *http.Request) {
		message = "404 page not found"
		w.Write([]byte(message))
	})

	server := http.Server{
		Addr:    ":8000",
		Handler: mux,
	}
	server.ListenAndServe()
}
