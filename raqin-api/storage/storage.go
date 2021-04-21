package storage

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var dbInstance *sql.DB

func Connect() error {

	dbURI := fmt.Sprintf(
		"user=%s password=%s host=%s dbname=%s sslmode=disable TimeZone=Asia/Muscat",
		"user", "password", "localhost", "raqin")

	var err error
	dbInstance, err = sql.Open("postgres", dbURI)
	if err != nil {
		panic(err)
	}

	if dbInstance.Ping() != nil {
		panic("Can not connect to DB")
	}

	return nil
}

func DBInstance() *sql.DB {
	return dbInstance
}
