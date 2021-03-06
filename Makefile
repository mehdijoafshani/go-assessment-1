test:
	go test --tags=exp ./...

run:
	go run main/main.go

dep:
	go mod tidy
	go mod download

build:
	go build ./...