package main

import (
	"context"
	"log"
)

const apiKey = ""

func main() {
	notion, err := NewClient("https://api.notion.com/v1", apiKey, nil)

	if err != nil {
		panic(err)
	}

	response, err := notion.GetMe(context.Background())

	if err != nil {
		panic(err)
	}

	// fmt.Println(response)

	log.Println(response)
}
