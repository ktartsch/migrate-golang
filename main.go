package main

import (
	"fmt"
	"github.com/ktartsch/migrate-golang/pkg/database/postgresql"
	"github.com/ktartsch/migrate-golang/pkg/model"
)

const dbHost = "db"
const port = 5432

func main() {

	fmt.Println("starting service")

	//config db

	dbCfg := postgresql.Config{
		Host:         dbHost,
		Port:         port,
		DatabaseName: "demo",
		Username:     "demo",
		Password:     "demo",
	}

	storage, err := postgresql.NewStorage(&dbCfg)

	if err != nil {
		panic(err)
	}

	persons, err := storage.Persons()

	if err != nil {
		panic(err)
	}

	fmt.Println("Sum persons: ", len(persons))

	err = storage.AddPerson(model.Person{
		FirstName: "John",
		LastName:  "Wayne",
	})

	if err != nil {
		panic(err)
	}

	persons, err = storage.Persons()

	if err != nil {
		panic(err)
	}

	fmt.Println("Sum persons: ", len(persons))
}
