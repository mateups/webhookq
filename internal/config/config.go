package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	APIListenAddr   string
	WorkerPollMs    int
	PostgresDsn     string
	ShutdownTimeout int
}

func Load() (Config, error) {
	pollMs := intFromEnv("WHQ_WORKER_POLL_MS", 1000)
	shutdownTimeout := intFromEnv("WHQ_SHUTDOWN_TIMEOUT_SEC", 10)

	cfg := Config{
		APIListenAddr:   stringFromEnv("WHQ_API_ADDR", ":8080"),
		WorkerPollMs:    pollMs,
		PostgresDsn:     stringFromEnv("WHQ_PG_DSN", ""),
		ShutdownTimeout: shutdownTimeout,
	}

	if cfg.WorkerPollMs <= 0 {
		return Config{}, fmt.Errorf("WHQ_WORKER_POLL_MS must be > 0")
	}
	if cfg.ShutdownTimeout <= 0 {
		return Config{}, fmt.Errorf("WHQ_SHUTDOWN_TIMEOUT_SEC must be > 0")
	}

	return cfg, nil
}

func stringFromEnv(key, fallback string) string {
	v := os.Getenv(key)
	if v == "" {
		return fallback
	}
	return v
}

func intFromEnv(key string, fallback int) int {
	v := os.Getenv(key)
	if v == "" {
		return fallback
	}
	n, err := strconv.Atoi(v)
	if err != nil {
		return fallback
	}
	return n
}
