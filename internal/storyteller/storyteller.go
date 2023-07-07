package storyteller

import (
	"encoding/json"
	"log"
	"os"
)

type Option struct {
	Text string `json:"text"`
	Next string `json:"arc"`
}

type Arc struct {
	Title   string   `json:"title,omitempty"`
	Story   []string `json:"story,omitempty"`
	Options []Option `json:"options"`
}

func NewStoryTeller(filePath string) (map[string]Arc, error) {

	var stories map[string]Arc

	// read file
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Println("error reading file")
		return nil, err
	}

	// parse into struct
	err = json.Unmarshal(data, &stories)
	if err != nil {
		log.Println("error unmarshalling json")
		return nil, err
	}

	return stories, nil
}
