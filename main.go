package main

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
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
	notion, err := NewClient("https://api.notion.com/v1", accessToken, nil)

	if err != nil {
		panic(err)
	}

	// response, err := notion.GetMe(context.Background())
	response, err := notion.GetRetrievePage(context.Background(), "185fdb9143bd42ab9dedca84f171afaf")

	if err != nil {
		panic(err)
	}

	log.Println(response)
}
