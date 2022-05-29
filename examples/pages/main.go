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

	response, err := client.GetRetrievePage(context.Background(), "762e592d92ca49abac6903a3c2f4c9d5")

	if err != nil {
		panic(err)
	}

	log.Println(response)

}
