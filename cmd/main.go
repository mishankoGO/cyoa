package main

import (
	"fmt"
	"github.com/mishankoGO/cyoa/internal/storyteller"
	"log"
)

func main() {
	filePath := "gopher.json"

	storyTeller, err := storyteller.NewStoryTeller(filePath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(storyTeller["debate"].Title)
}
