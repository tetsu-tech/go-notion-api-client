package main

import (
	"context"
	"log"
	"os"
)

func main() {
	accessToken := os.Getenv("GO_NOTION_ACCESS_TOKEN")
	notion, err := NewClient("https://api.notion.com/v1", accessToken, nil)

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
