// package main

// import (
// 	"net/http"
// )

// func main() {
// 	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
// 		var message string
// 		if r.URL.Path == "/" {
// 			message = "root web"
// 			w.Write([]byte(message))
// 		} else if r.URL.Path == "/products" {
// 			if r.Method == "GET" {
// 				message = "get products"
// 				w.Write([]byte(message))
// 			} else if r.Method == "DELETE" {
// 				message = "delete products"
// 				w.Write([]byte(message))
// 			}
// 		} else if r.URL.Path == "/users" {
// 			message = "root web"
// 			w.Write([]byte(message))
// 		} else {
// 			message = "404 page not found"
// 			w.Write([]byte(message))
// 		}
// 	}

// 	server := http.Server{
// 		Addr:    ":8000",
// 		Handler: handler,
// 	}
// 	server.ListenAndServe()
// }
