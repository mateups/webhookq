package main

import (
	"log"
	"os"

	"webhooq/internal/app/worker"
	"webhooq/internal/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("load config: %v", err)
	}

	if err := worker.Run(cfg, os.Stdout); err != nil {
		log.Fatalf("worker exited with error: %v", err)
	}
}
