package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
)

var query *Queries
var connection *sql.DB
var err error

func Init() {
	connection, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	query = New(connection)
}

func GetConnection() *sql.DB {
	return connection
}

func GetQuery() *Queries {
	return query
}
