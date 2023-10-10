package config

import (
	"log"

	"github.com/hamdiBouhani/my-playground-project/storage"
	"github.com/hamdiBouhani/my-playground-project/storage/pg"
)

func NewDBlient() (storage.Storage, error) {
	db := pg.NewDBConn()
	err := db.CreateConnection()
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = db.Migrate()
	if err != nil {
		log.Fatalf(err.Error())
	}
	return db, nil
}
