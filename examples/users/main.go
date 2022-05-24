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
	client, err := notion.NewClient(accessToken, nil)

	if err != nil {
		panic(err)
	}

	response, err := client.GetMe(context.Background())

	if err != nil {
		panic(err)
	}

	log.Println(response)

}
