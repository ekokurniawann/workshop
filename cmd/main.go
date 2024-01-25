package main

import (
	"fmt"
	"net/http"
	"workshop/internal"
	"workshop/models"

	"github.com/gorilla/mux"
)

func main() {
	models.InitDB()

	r := mux.NewRouter()

	r.HandleFunc("/test", internal.Test).Methods("GET")
	r.HandleFunc("/create", internal.Create).Methods("POST")
	r.HandleFunc("/people", internal.GetPeople).Methods("GET")
	r.HandleFunc("/person/{id}", internal.GetPerson).Methods("GET")
	r.HandleFunc("/update-person/{id}", internal.UpdatePerson).Methods("PUT")
	r.HandleFunc("/delete-person/{id}", internal.DeletePerson).Methods("DELETE")

	address := ":8080"
	fmt.Printf("server started at %s\n", address)
	http.ListenAndServe(address, r)
}
