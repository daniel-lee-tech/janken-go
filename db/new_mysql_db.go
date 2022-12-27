package db

import (
	"database/sql"
	"log"
)

func NewDbConnection() (*sql.DB, error) {
	db, err := sql.Open("mysql", "dev_user:dev_password@tcp(127.0.0.1:3306)/janken_dev")
	if err != nil {
		log.Fatal(err)
	}

	var version string

	err2 := db.QueryRow("SELECT VERSION()").Scan(&version)

	if err2 != nil {
		log.Fatal(err2)
	}

	return db, nil
}
