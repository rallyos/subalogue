package db

import (
	"database/sql"
	"log"
	"os"
)

var Query *Queries

func Init() *sql.DB {
	conn, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	Query = New(conn)
	return conn
}
