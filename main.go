package main

import (
	"fmt"
	"github.com/ktartsch/migrate-golang/pkg/database/postgresql"
	"github.com/ktartsch/migrate-golang/pkg/model"
)

func main() {

	fmt.Println("starting service")

	//config db

	dbCfg := postgresql.Config{
		Host:         "db",
		Port:         5432,
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
