package db

import (
	"database/sql"
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type postgresDatabase struct {
	dsn string
	db  *sql.DB
}

func NewPostgresDatabase(dsn string) Database {
	return &postgresDatabase{dsn: dsn, db: nil}
}

// Close implements [Database].
func (p *postgresDatabase) Close() error {
	if p.db == nil {
		return nil
	}

	return p.db.Close()
}

// GetInstance implements [Database].
func (p *postgresDatabase) GetInstance() (*sql.DB, error) {
	if p.db == nil {
		return nil, fmt.Errorf("database is not open")
	}

	return p.db, nil
}

// MigrateOnStartup implements [Database].
func (p *postgresDatabase) MigrateOnStartup() error {
	if p.db == nil {
		return fmt.Errorf("database is not open")
	}

	driver, err := postgres.WithInstance(p.db, &postgres.Config{})

	if err != nil {
		return err
	}

	_, currentFile, _, ok := runtime.Caller(0)
	if !ok {
		return fmt.Errorf("failed to resolve migrations path")
	}

	migrationsPath := filepath.Join(filepath.Dir(currentFile), "migrations")

	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", migrationsPath),
		"postgres", driver)

	if err != nil {
		return err
	}

	if err := m.Up(); err != nil {
		if err != migrate.ErrNoChange {
			return err
		}
		fmt.Println("No database migrations to apply")
	}

	return nil
}

// Open implements [Database].
func (p *postgresDatabase) Open(dsn string) error {
	db, err := sql.Open("postgres", dsn)

	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	p.db = db

	return nil
}
