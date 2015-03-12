package main

import (
	"database/sql"
	"log"
	"net/url"

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

func createRedirect(vals url.Values) bool {
	//add validation
	res, err := db.Exec("INSERT INTO redirects (host, rules) VALUES($1, $2);", vals["host"][0], vals["rules"][0])
	if c, _ := res.RowsAffected(); c == 0 {
		log.Println(err)
		return false
	}
	return true
}

func lookup(host string) string {
	dest := ""
	err := db.Get(&dest, "SELECT rules FROM redirects WHERE host = $1", host)
	if !noRows(err) {
		log.Println(err)
	}
	return dest
}
