package main

import (
	"context"
	"database/sql"
	"github.com/dmralev/subalogue/repositories"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func main() {

	conn, err := sql.Open("postgres", "user=subalogue password=subalogue dbname=subalogue_development host=db sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	query := db.New(conn)
	ctx := context.Background()
	harcoded_user_id := sql.NullInt32{Int32: 1, Valid: true}

	r := gin.Default()

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "Ping")
	})

	r.GET("me/subscriptions", func(c *gin.Context) {
		// TODO Check if null and flip the Valid key when true
		subscriptions, err := query.ListUserSubscriptions(ctx, harcoded_user_id)
		if err != nil {
			log.Fatal(err)
		}

		c.JSON(200, subscriptions)
	})

	r.POST("me/subscriptions", func(c *gin.Context) {
		var subscription_json db.CreateUserSubscriptionsParams
		subscription_json.UserID = harcoded_user_id

		if err := c.ShouldBindJSON(&subscription_json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		query.CreateUserSubscriptions(ctx, subscription_json)

		c.Status(http.StatusCreated)
	})

	r.Run()
}
