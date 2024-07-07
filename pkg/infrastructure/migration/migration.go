package migration

import (
	"fmt"
	"log"

	"github.com/AdiPP/go-marketplace/pkg/infrastructure/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type PostgresMigration struct {
	config config.Database
}

func NewPostgresMigration(config config.Database) *PostgresMigration {
	return &PostgresMigration{
		config: config,
	}
}

func (p *PostgresMigration) Migrate() {
	m, err := migrate.New(
		"file://db/migrations",
		fmt.Sprintf(
			"postgres://%s:%s@%s:%s/%s?sslmode=disable&search_path=%s",
			p.config.DatabaseUser,
			p.config.DatabasePassword,
			p.config.DatabaseHost,
			p.config.DatabasePort,
			p.config.DatabaseName,
			p.config.DatabaseSchema,
		),
	)
	if err != nil {
		log.Fatal("Could not create migrate instance: ", err)
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Fatal("Migration failed: ", err)
	}

	log.Println("Migration succeed")
}
