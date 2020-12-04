package main

import (
	"database/sql"
	"fmt"
	"log"
)

func connect() *sql.DB {
	const (
		hostname      = "localhost"
		host_port     = 6969
		username      = "postgres"
		password      = "T3p4tBTPNS"
		database_name = "cobagolang"
	)
	pg_con_string := fmt.Sprintf("port=%d host=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host_port, hostname, username, password, database_name)
	db, err := sql.Open("postgres", pg_con_string)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
