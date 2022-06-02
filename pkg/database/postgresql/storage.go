package postgresql

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/ktartsch/migrate-golang/pkg/model"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type Storage struct {
	db *sqlx.DB
}

// Config configures the connection to the PostgreSQL database.
type Config struct {
	Host         string
	Port         int `default:"5432"`
	DatabaseName string
	Username     string
	Password     string
}

//NewStorage creates a ready to use *Storage.
func NewStorage(cfg *Config) (*Storage, error) {

	str := &Storage{}

	dbConn, err := str.connect(cfg)

	if err != nil {
		return nil, err
	}

	err = migrateDB(dbConn, cfg.DatabaseName)
	if err != nil {
		return nil, err
	}

	str.db = dbConn

	return str, nil
}

func (str *Storage) connect(cfg *Config) (*sqlx.DB, error) {

	connectionString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host,
		cfg.Port,
		cfg.Username,
		cfg.Password,
		cfg.DatabaseName,
	)
	return sqlx.Connect("pgx", connectionString)
}

// Persons fetches all persons from storage without any condition.
func (str *Storage) AddPerson(p model.Person) error {

	stmt := `INSERT INTO person (first_name, last_name) VALUES($1, $2);`

	_, err := str.db.Exec(stmt, p.FirstName, p.LastName)
	if err != nil {
		return err
	}

	return nil

}

// Persons fetches all persons from storage without any condition.
func (str *Storage) Persons() ([]*model.Person, error) {

	stmt := `SELECT id, first_name, last_name FROM person;`

	persons := make([]*model.Person, 0)

	err := str.db.Select(&persons, stmt)

	if err != nil {
		return nil, err
	}

	return persons, nil

}
