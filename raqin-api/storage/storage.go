package storage

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var dbInstance *sql.DB

func Connect() error {

	dbURI := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s",
		"root", "password", "localhost", "raqin")

	var err error
	dbInstance, err = sql.Open("mysql", dbURI)
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
