package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

const PORT = ":8080"

type Pegawai struct {
	ID     int
	Name   string
	Age    int
	Divisi string
}

var karyawan = []Pegawai{}

func init() {
	karyawan = []Pegawai{
		{ID: 1, Name: "Jono", Age: 22, Divisi: "HRD"},
		{ID: 1, Name: "Jupri", Age: 29, Divisi: "DIREKTUR"},
		{ID: 1, Name: "Jini", Age: 19, Divisi: "OPS"},
		{ID: 1, Name: "Jack", Age: 40, Divisi: "IT"},
	}
}

func main() {
	http.HandleFunc("/pegawai", dataPegawai)
	fmt.Println("Application Runing on Port ", PORT)
	log.Fatal(http.ListenAndServe(PORT, nil))
}

/**
*  Request form : name , age
 */
func dataPegawai(w http.ResponseWriter, r *http.Request) {
	//header response
	w.Header().Set("Content-Type", "application/json")
	if r.Method == http.MethodPost {
		name := r.FormValue("name")
		divisi := r.FormValue("divisi")
		age, err := strconv.Atoi(r.FormValue("age"))
		if err != nil {
			http.Error(w, "Invalid age", http.StatusBadRequest)
			return
		}
		newKaryawan := Pegawai{
			ID:     len(karyawan) + 1,
			Name:   name,
			Age:    age,
			Divisi: divisi,
		}
		fmt.Println(" New Karyawan : ", newKaryawan)
		karyawan = append(karyawan, newKaryawan)
		json.NewEncoder(w).Encode(karyawan)
		return
	}

	if r.Method == http.MethodGet {
		json.NewEncoder(w).Encode(karyawan)
		return
	}

	http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)

}
