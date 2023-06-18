package main

import (
	"fmt"
	"go-api-rest/database"
	"go-api-rest/models"
	"go-api-rest/routes"
)

func main() {
	models.Personalities = []models.Personality {
		{
			Id: 1,
			Name: "Person A",
			History: "Lorem ipsum dolor sit amet ...",
		},
		{
			Id: 2,
			Name: "Person B",
			History: "Lorem ipsum dolor sit amet ...",
		},
	}

	database.StablishConnection()
	fmt.Println("Starting server with Go")
	routes.HandleRequest()
}
