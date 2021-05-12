package main

import (
	"raqin-api/api"
	"raqin-api/app"
	"raqin-api/storage"
)

func main() {

	app.Init()
	err := storage.Connect()
	if err != nil {
		panic(err)
	}
	api.Init()
}
