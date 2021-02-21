test:
	go test --tags=exp ./...

run:
	go run main/main.go

depen:
	go mod download

build:
	go build ./...