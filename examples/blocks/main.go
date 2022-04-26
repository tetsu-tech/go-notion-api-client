package main

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/tetsu-tech/go-notion-api-client"
)

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
}

func main() {
	loadEnv()

	accessToken := os.Getenv("GO_NOTION_ACCESS_TOKEN")
	client, err := notion.NewClient("https://api.notion.com/v1", accessToken, nil)

	if err != nil {
		panic(err)
	}

	if len(os.Args) <= 1 {
		log.Panic("blockID is not specified")
	}

	blockID := os.Args[1]

	response, err := client.RetrieveBlock(context.Background(), blockID)

	if err != nil {
		panic(err)
	}

	log.Println(response)
}
