package database

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // mysql driver
	"github.com/jmoiron/sqlx"
)

func New(user, pass, host, name string, port int) (*sqlx.DB, func(), error) {
	db, err := sqlx.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", user, pass, host, port, name))
	if err != nil {
		return nil, nil, err
	}

	cleanup := func() {
		if err := db.Close(); err != nil {
			log.Print("failed to close database connection")
		} else {
			log.Print("successfully closed database connection")
		}
	}

	return db, cleanup, nil
}
