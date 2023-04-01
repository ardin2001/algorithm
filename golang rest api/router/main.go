package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func rootWeb(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	reqHeader := r.Header.Get("content-type")
	w.Header().Add("Created-by", "Ardin Nugraha")
	message := "root file " + r.URL.Path + " : " + reqHeader
	w.Write([]byte(message))
}

func getProducts(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := params.ByName("id")
	name := params.ByName("name")
	message := "get products by : " + id + " : " + name
	w.Write([]byte(message))
}
func deleteProduct(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := params.ByName("id")
	if id == "1" {
		w.WriteHeader(401)
		return
	}
	name := params.ByName("name")
	message := "get products by : " + id + " : " + name
	w.Write([]byte(message))
}
func getProductDetail(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	details := params.ByName("details")
	message := "get products by : " + details
	w.Write([]byte(message))
}

func main() {
	router := httprouter.New()
	router.GET("/", rootWeb)
	router.GET("/products/:id/:name", getProducts)
	router.GET("/users/*details", getProductDetail)
	router.DELETE("/users/:id", deleteProduct)

	server := http.Server{
		Addr:    ":8000",
		Handler: router,
	}
	server.ListenAndServe()
}
