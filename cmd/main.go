package main

import (
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

	controller := controllers.NewController(storyTeller)
	router := controller.Route()

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	log.Println("Listening...")
	server.ListenAndServe()
	//cli := cli2.NewCli(storyTeller)
	//err = cli.Game()
	//if err != nil {
	//	log.Fatal(err)
	//}
}
