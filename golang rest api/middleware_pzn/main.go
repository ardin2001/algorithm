package main

import (
	"fmt"
	"net/http"
)

type LogMiddleware struct {
	Handler http.Handler
}

func (middleware *LogMiddleware) serveHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("before check middleware")
	middleware.Handler.ServeHTTP(w, r)
	fmt.Println("after check middleware")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("handler executed")
		fmt.Fprint(w, "hello middleware")
	})

	logMiddleware := &LogMiddleware{
		Handler: mux,
	}
	server := http.Server{
		Addr:    "8000",
		Handler: logMiddleware,
	}
	fmt.Println("Server is running in port 8000")
	server.ListenAndServe()

}
