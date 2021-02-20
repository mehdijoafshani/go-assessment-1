package main

import (
	"github.com/mehdijoafshani/go-assessment-1/internal/config"

	// to run init and setup logger
	_ "github.com/mehdijoafshani/go-assessment-1/internal/logger"
)

func main() {
	config.Setup()
	//TODO start grpc/rest server
}
