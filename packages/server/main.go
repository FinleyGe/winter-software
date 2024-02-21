package main

import (
	"SelectionSystem/database"
	"SelectionSystem/router"
)

func main() {
	database.Init()

	r := router.SetRouter()

	r.Run(":8686")
}
