package main

import (
	"fmt"
	"net/http"
)

// Fungi Log yang berguna sebagai middleware
func log(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Ini dari middleware Log....\n")
		next.ServeHTTP(w, r)
	})
}

// Fungsi GetMahasiswa untuk mengampilkan text string di browser
func GetMahasiswa(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ini dari function GetMahasiswa()"))
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ini dari function GetUsers()"))
}

func main() {

	// konfigurasi server
	server := &http.Server{
		Addr: ":8000",
	}
	// routing
	http.Handle("/", log(GetMahasiswa))
	// http.Handle("/users", log(http.HandlerFunc(GetUsers)))

	// jalankan server
	server.ListenAndServe()
}
