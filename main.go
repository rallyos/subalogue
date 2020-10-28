package main

import (
	"subalogue/app"
)

func main() {
	server := app.Server{}
	server.Initialize()
	server.Run()
}
