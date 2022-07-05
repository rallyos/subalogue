package main

import (
	"github.com/shiftingphotons/subalogue/app"
)

func main() {
	server := app.Server{}
	server.Initialize()
	server.Run()
}
