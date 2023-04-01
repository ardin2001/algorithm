package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Barang struct {
	Nama        string `json:"nama"`
	Harga       int    `json:"harga"`
	JenisBarang string `json:"jenis_barang"`
}

var data = []Barang{
	{"AAA", 67000, "Makanan"},
	{"BBB", 67000, "Minuman"},
	{"CCC", 67000, "Makanan"},
	{"DDD", 67000, "Minuman"},
}

func main() {
	http.HandleFunc("/barang", GetBarang)
	http.HandleFunc("/barang-id", GetBarangByNama)

	fmt.Println("server is running in port : 8000")
	http.ListenAndServe(":8000", nil)
}

func GetBarang(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	if r.Method == "GET" {
		res, err := json.Marshal(data)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(res)
		return
	} else if r.Method == "DELETE" {
		message := "barang berhasil dihapus" + r.URL.RawQuery + " : " + r.URL.Path
		w.Write([]byte(message))
		return
	}

	http.Error(w, "400 status bad request", http.StatusBadRequest)
}

func GetBarangByNama(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	if r.Method == "GET" {
		nama := r.FormValue("nama123")

		for _, v := range data {
			if v.Nama == nama {
				result, err := json.Marshal(v)

				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				w.Write(result)
				return
			}
		}
	}

	http.Error(w, "400 status bad request", http.StatusBadRequest)
}
