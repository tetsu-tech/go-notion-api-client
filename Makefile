go_build:
	go build && ./go-notion-api-client

go_run-%:
	go run main.go client.go ${@:go_run-%=%}.go
