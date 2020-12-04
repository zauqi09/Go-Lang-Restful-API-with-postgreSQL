package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func getUsers(w http.ResponseWriter, r *http.Request) {

	var response Response
	arruser := getDBUser()

	response.Status = 200
	response.Message = "Getting all users success"
	response.Data = arruser

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func getAnUser(w http.ResponseWriter, r *http.Request) {
	var response Response
	status := 404
	message := "User not found"

	params := mux.Vars(r)
	id := params["id"]

	arruser := getAnDBUser(id)
	if len(arruser) != 0 {
		status = 200
		message = "Success getting an User"
	}
	response.Status = status
	response.Message = message
	response.Data = arruser

	log.Print(response)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func insertUser(w http.ResponseWriter, r *http.Request) {
	var response Response
	var users Users
	var arruser []Users
	status := 200
	message := "Insert user success"

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(body, &users)
	if err != nil {

		panic(err)
	}
	if users.FirstName != "" || users.LastName != "" {
		addDBUser(users)
	} else {
		status = 400
		message = "Bad Request"
	}

	arruser = append(arruser, users)

	response.Status = status
	response.Message = message
	response.Data = arruser
	log.Print(response)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	var response Response
	status := 404
	message := "User not found"
	params := mux.Vars(r)
	id := params["id"]

	arruser := getAnDBUser(id)
	if len(arruser) != 0 {
		delDBUser(id)
		status = 200
		message = "Delete user success"
	}
	response.Data = arruser
	response.Status = status
	response.Message = message

	log.Print(response)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func updateUser(w http.ResponseWriter, r *http.Request) {
	var response Response
	var users Users
	status := 404
	message := "User not found"
	params := mux.Vars(r)
	id := params["id"]

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(body, &users)

	arruser := getAnDBUser(id)
	if len(arruser) != 0 {
		if users.FirstName != "" || users.LastName != "" {
			updateDBUser(id, users)
			status = 200
			message = "Update user success"
		} else {
			status = 400
			message = "Bad Request"
		}
	}
	response.Data = arruser
	response.Status = status
	response.Message = message

	log.Print(response)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}
