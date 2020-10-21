package db

import (
	"database/sql"
	"log"
	"os"
)

var Query *Queries

func init() {
	conn, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	Query = New(conn)
}
