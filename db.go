package main

import (
	"database/sql"

	_ "github.com/lib/pq"

	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

const MAX_DB_CONNECTIONS = 10

func connectDB() {
	db = sqlx.MustOpen("postgres", "host=/var/run/postgresql dbname=redirector_development sslmode=disable")
	db.SetMaxOpenConns(MAX_DB_CONNECTIONS)
}

func noRows(err error) bool {
	return sql.ErrNoRows == err
}
