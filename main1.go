package main

import (
	"restful-api/config"
	"restful-api/route"
)

func main() {
	config.InitDB()
	e := route.New()
	e.Start("127.0.0.1:8000")
}
