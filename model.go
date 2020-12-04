package main

import (
	"log"
)

type Users struct {
	Id        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Users
}

func addDBUser(users Users) {
	db := connect()
	defer db.Close()

	_, err := db.Exec("INSERT INTO public.person (first_name, last_name) values ($1,$2)",
		users.FirstName,
		users.LastName,
	)

	if err != nil {
		log.Print(err)
	}
}

func getDBUser() []Users {
	var users Users
	var arruser []Users
	db := connect()
	defer db.Close()

	rows, err := db.Query("Select id,first_name,last_name from public.person")
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&users.Id, &users.FirstName, &users.LastName); err != nil {
			log.Fatal(err.Error())

		} else {
			arruser = append(arruser, users)
		}
	}
	return arruser
}

func delDBUser(id string) {
	db := connect()
	defer db.Close()

	_, err := db.Exec("DELETE from public.person WHERE id=$1", id)
	if err != nil {
		log.Print(err)
	}
}

func getAnDBUser(id string) []Users {
	var users Users
	var arruser []Users
	db := connect()
	defer db.Close()

	rows, err := db.Query("Select id,first_name,last_name from public.person where id=$1", id)
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&users.Id, &users.FirstName, &users.LastName); err != nil {
			log.Fatal(err.Error())

		} else {
			arruser = append(arruser, users)
		}
	}
	return arruser
}

func updateDBUser(id string, users Users) {
	db := connect()
	defer db.Close()

	_, err := db.Query("UPDATE public.person SET first_name=$1, last_name=$2 where id=$3", users.FirstName, users.LastName, id)
	if err != nil {
		log.Print(err)
	}
}
