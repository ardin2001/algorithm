package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// Mahasiswa
type Mahasiswa struct {
	ID   int    `json:"id"`
	NIM  int    `json:"nim"`
	Name string `json:"name"`
}

// NewMahasiswa
var NewMahasiswa = []Mahasiswa{
	{
		ID:   1,
		NIM:  123454,
		Name: "Didik Prabowo",
	},
	{
		ID:   2,
		NIM:  923454,
		Name: "Joni Gunawan",
	},
	{
		ID:   3,
		NIM:  923454,
		Name: "Muhammad Irwan",
	},
}

// GetMahasiswa
func GetMahasiswa(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		datamahasiswa, err := json.Marshal(NewMahasiswa)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(datamahasiswa)
		return
	}

	http.Error(w, "hayo mau ngapain", http.StatusNotFound)
}

// PostMahasiswa
func PostMahasiswa(w http.ResponseWriter, r *http.Request) {
	fmt.Println("masuk post")
	w.Header().Set("Content-Type", "application/json")
	var Mhs Mahasiswa
	if r.Method == "POST" {
		if r.Header.Get("Content-Type") == "application/json" {
			// parse dari json
			decodeJSON := json.NewDecoder(r.Body)
			if err := decodeJSON.Decode(&Mhs); err != nil {
				log.Fatal(err)
			}
		} else {
			// parse dari form
			getID := r.PostFormValue("id")
			id, _ := strconv.Atoi(getID)
			getNim := r.PostFormValue("nim")
			nim, _ := strconv.Atoi(getNim)
			name := r.PostFormValue("name")
			Mhs = Mahasiswa{
				ID:   id,
				NIM:  nim,
				Name: name,
			}
		}

		NewMahasiswa = append(NewMahasiswa, Mhs)
		dataMahasiswa, _ := json.Marshal(Mhs) // to byte
		w.Write(dataMahasiswa)                // cetak di browser
		return
	}

	http.Error(w, "hayo mau ngapain", http.StatusNotFound)
}

func PutMahasiswa(w http.ResponseWriter, r *http.Request) {
	fmt.Println("masuk put")
	id := r.URL.Query().Get("id")
	w.Header().Set("Content-Type", "application/json")
	var Mhs Mahasiswa
	if r.Method == "PUT" {
		if r.Header.Get("Content-Type") == "application/json" {
			fmt.Println("masuk json")
			// parse dari json
			decodeJSON := json.NewDecoder(r.Body)
			if err := decodeJSON.Decode(&Mhs); err != nil {
				log.Fatal(err)
			}

		} else {
			// parse dari form
			getID := r.PostFormValue("id")
			id, _ := strconv.Atoi(getID)
			getNim := r.PostFormValue("nim")
			nim, _ := strconv.Atoi(getNim)
			name := r.PostFormValue("name")
			Mhs = Mahasiswa{
				ID:   id,
				NIM:  nim,
				Name: name,
			}
		}

		for i, data := range NewMahasiswa {
			val, _ := strconv.Atoi(id)
			if val == data.ID {
				fmt.Println(data)
				NewMahasiswa[i].ID = val
				NewMahasiswa[i].NIM = Mhs.NIM
				NewMahasiswa[i].Name = Mhs.Name
				fmt.Println(val)
			}
		}

		fmt.Println(NewMahasiswa)
		dataMahasiswa, _ := json.Marshal(Mhs)
		w.Write(dataMahasiswa)
		return
	}

	http.Error(w, "hayo mau ngapain", http.StatusNotFound)
}

func DelMahasiswa(w http.ResponseWriter, r *http.Request) {
	fmt.Println("masuk delete")
	id := r.URL.Query().Get("id")
	w.Header().Set("Content-Type", "application/json")
	var Mhs *Mahasiswa
	if r.Method == "DELETE" {
		newUsers := make([]Mahasiswa, 0)
		for i, data := range NewMahasiswa {
			val, _ := strconv.Atoi(id)
			if val == data.ID {
				Mhs = &data
				fmt.Println(Mhs)
				newUsers = append(newUsers, NewMahasiswa[:i]...)
				newUsers = append(newUsers, NewMahasiswa[i+1:]...)

				NewMahasiswa = newUsers
			}
		}

		fmt.Println(&NewMahasiswa)
		dataMahasiswa, _ := json.Marshal(&Mhs)
		w.Write(dataMahasiswa)
		return
	}

	http.Error(w, "hayo mau ngapain", http.StatusNotFound)
}

func main() {
	http.HandleFunc("/mahasiswa", GetMahasiswa)
	http.HandleFunc("/post-mahasiswa", PostMahasiswa)
	http.HandleFunc("/put-mahasiswa", PutMahasiswa)
	http.HandleFunc("/del-mahasiswa", DelMahasiswa)
	fmt.Println("server running...")
	if err := http.ListenAndServe(":7000", nil); err != nil {
		log.Fatal(err)
	}
}
