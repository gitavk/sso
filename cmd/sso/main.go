package main

import (
	"fmt"

	"github.com/gitavk/sso.v1/internal/config"
)

func main() {
	cfg := config.MustLoad()
	fmt.Printf("Loaded config: %+v\n", cfg)
	// TODO: logger
	// TODO: init app
	// TODO: run gRPC server
}
