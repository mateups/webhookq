package main

import (
	"log"
	"os"

	"webhooq/internal/app/api"
	"webhooq/internal/config"
	"webhooq/internal/db"
	"webhooq/internal/targets"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("load config: %v", err)
	}

	database := db.NewPostgresDatabase(cfg.PostgresDsn)
	if err := database.Open(cfg.PostgresDsn); err != nil {
		log.Fatalf("open database: %v", err)
	}
	defer database.Close()

	if err := database.MigrateOnStartup(); err != nil {
		log.Fatalf("migrate database: %v", err)
	}

	databaseInstance, err := database.GetInstance()
	if err != nil {
		log.Fatalf("get database instance: %v", err)
	}

	targetsRepository := targets.NewPostgresRepository(databaseInstance)
	targetsService := targets.NewService(targetsRepository)
	targetsHandler := api.NewTargetsHandler(targetsService)

	if err := api.Run(cfg, os.Stdout, targetsHandler); err != nil {
		log.Fatalf("api exited with error: %v", err)
	}
}
