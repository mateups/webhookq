package db

import "database/sql"

type Database interface {
	Close() error
	GetInstance() (*sql.DB, error)
	MigrateOnStartup() error
	Open(dsn string) error
}
