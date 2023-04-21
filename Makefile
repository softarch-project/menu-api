server:
	go run main.go

deps:
	go mod download

docs:
	swag init -g httpserver/httpserver.go