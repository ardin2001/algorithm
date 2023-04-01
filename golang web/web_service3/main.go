package main

import (
	"fmt"
	"net/http"
)

// func middleware(next func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
func middleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(rWriter http.ResponseWriter, request *http.Request) {
		// Lakukan sesuatu di middleware sebelum memanggil handler selanjutnya
		fmt.Println("masuk middleware")
		next(rWriter, request)
		// Lakukan sesuatu di middleware setelah panggilan ke handler selesai
	})
}

func rootWeb(w http.ResponseWriter, r *http.Request) {
	message := "root web"
	fmt.Println("root web")
	w.Write([]byte(message))
}

func main() {
	mux := http.NewServeMux()
	var message string
	mux.HandleFunc("/", middleware(rootWeb))
	mux.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			message = "get products"
			w.Write([]byte(message))
		} else if r.Method == "DELETE" {
			message = "delete products"
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
