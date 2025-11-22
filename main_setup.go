package main

import (
	"database/sql"
	"errors"
	"os"

	"github.com/adamjames870/tripstats/internal/database"
)

func (state *apiState) LoadState() error {
	db, errDb := loadDb()
	if errDb != nil {
		return errDb
	}
	state.db = database.New(db)
	state.secret_string = loadSecretString()
	return nil
}

func loadDb() (*sql.DB, error) {
	dbUrl := os.Getenv("DB_URL")
	db, errDb := sql.Open("postgres", dbUrl)
	if errDb != nil {
		return nil, errors.New("Unable to connect to database" + errDb.Error())
	}
	return db, nil
}

func loadSecretString() string {
	return os.Getenv("SECRET")
}
