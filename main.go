package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/user", getUsers).Methods("GET")
	router.HandleFunc("/user/{id}", getAnUser).Methods("GET")
	router.HandleFunc("/user", insertUser).Methods("POST")
	router.HandleFunc("/user/{id}", deleteUser).Methods("DELETE")
	router.HandleFunc("/user/{id}", updateUser).Methods("PUT")
	http.Handle("/", router)
	fmt.Println("Connected to port 1234")
	log.Fatal(http.ListenAndServe(":1234", router))

}
