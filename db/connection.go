package db

import (
	"database/sql"
	"log"
)

var Query *Queries

func Init() {
	conn, err := sql.Open("postgres", "user=subalogue password=subalogue dbname=subalogue_development host=db sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	Query = New(conn)
}
