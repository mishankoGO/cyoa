package main

import (
	"fmt"
	"github.com/mishankoGO/cyoa/internal/controllers"
	"github.com/mishankoGO/cyoa/internal/storyteller"
	"log"
	"net/http"
)

func main() {
	filePath := "gopher.json"

	storyTeller, err := storyteller.NewStoryTeller(filePath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(storyTeller["debate"].Options[0].Next)

	controller := controllers.NewController(storyTeller)
	router := controller.Route()

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	log.Println("Listening...")
	server.ListenAndServe()
}
