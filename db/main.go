package db

import (
	"database/sql"
	"fmt"
)

type ConnectionFunc func(uri string) *sql.DB

//ConnectorDB is DB conneciton as Clousure
func ConnectorDB(driver string) ConnectionFunc {

	return func(uri string) *sql.DB {
		conn, err := sql.Open(driver, uri)
		if err != nil {
			fmt.Println(err)
		}

		return conn
	}
}
