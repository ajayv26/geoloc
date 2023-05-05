package dbmigrate

import (
	"database/sql"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrateUp(dbSource string) {
	db, err := sql.Open("postgres", dbSource)
	if err != nil {
		log.Fatal("error when trying to connect:", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal("db instance is nil:", err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://sql/migrations", "postgres", driver)
	if err != nil {
		log.Fatal("migration files not found:", err)
	}

	// migrate up
	upErr := m.Up()
	if upErr != nil {
		if upErr == migrate.ErrNoChange {
			log.Println("migration is up to date:", upErr.Error())
		} else {
			log.Fatal("migration up error:", upErr.Error())
		}
	}

	if upErr == nil {
		log.Println("database migrated up successfully")
	}
}
