package main

import (
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	"subalogue/app"
)

func main() {
	server := app.Server{}
	server.Initialize()
	server.Run()
}
