package main

import (
	"database/sql"

	_ "github.com/lib/pq"

	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func connectDB() {
	db = sqlx.MustOpen("postgres", config.DB)
	db.SetMaxOpenConns(config.MaxDBConnections)
}

func noRows(err error) bool {
	return sql.ErrNoRows == err
}
