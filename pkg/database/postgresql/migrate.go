package postgresql

import (
	"embed"
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/pgx"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

//go:embed migrations
var fs embed.FS

const sourceName = "iofs"

func migrateDB(db *sqlx.DB, databaseName string) error {

	d, err := iofs.New(fs, "migrations")
	if err != nil {
		return err
	}

	driver, err := pgx.WithInstance(db.DB, &pgx.Config{
		MultiStatementEnabled: true,
	})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithInstance(sourceName, d, databaseName, driver)

	if err != nil {
		return fmt.Errorf("failed to initialize migrate instance, %w", err)
	}

	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	migrationVersion, dirty, _ := m.Version()

	fmt.Printf("migration version: %v isDirty: %v \n", migrationVersion, dirty)

	return nil
}
