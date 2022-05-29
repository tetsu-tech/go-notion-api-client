go_build:
	go build && ./go-notion-api-client

go_run-%:
	go run ./examples/${@:go_run-%=%}/main.go
