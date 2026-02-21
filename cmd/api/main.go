package main

import (
	"log"
	"os"

	"webhooq/internal/app/api"
	"webhooq/internal/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("load config: %v", err)
	}

	if err := api.Run(cfg, os.Stdout); err != nil {
		log.Fatalf("api exited with error: %v", err)
	}
}
