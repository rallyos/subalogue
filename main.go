package main

import (
	"context"
	"database/sql"
	"github.com/dmralev/subalogue/repositories"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
)

func main() {

	conn, err := sql.Open("postgres", "user=subalogue password=subalogue dbname=subalogue_development host=db sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	db := db.New(conn)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("me/subscriptions", func(c *gin.Context) {
		// TODO Check if null and flip the Valid key when true
		harcoded_user_id := sql.NullInt32{Int32: 1, Valid: true}
		subscriptions, err := db.ListUserSubscriptions(context.Background(), harcoded_user_id)
		if err != nil {
			log.Fatal(err)
		}

		c.JSON(200, subscriptions)
	})

	r.Run()
}
