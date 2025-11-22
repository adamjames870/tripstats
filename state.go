package main

import (
	"net/http"

	"github.com/adamjames870/tripstats/internal/database"
)

type apiState struct {
	db            *database.Queries
	mux           *http.ServeMux
	secret_string string
	tripApiKey    string
}
