package database

import (
	"embed"
	"os"
	"path/filepath"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/secretium/secretium/internal/config"
	"github.com/secretium/secretium/internal/constants"
)

//go:embed sql_queries/*
var sqlQueries embed.FS

// Database contains DB connection and other dependencies for application.
type Database struct {
	Connection *sqlx.DB
	SQLQueries embed.FS
}

// New returns a new instance of DB connection.
func New(c *config.Config) (*Database, error) {
	// Create a folder for the SQLite DB.
	_ = os.Mkdir(constants.ConstConfigSQLitePath, os.ModePerm)

	// Connect to the SQLite DB.
	connection, err := sqlx.Connect("sqlite3", filepath.Join(constants.ConstConfigSQLitePath, "db.sqlite3"))
	if err != nil {
		return nil, err
	}

	return &Database{
		Connection: connection,
		SQLQueries: sqlQueries,
	}, nil
}
