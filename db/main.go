package db

import "database/sql"

type ConnectionFunc func(uri string) (*sql.DB, error)

//ConnectorDB is DB conneciton as Clousure
func ConnectorDB(driver string) ConnectionFunc {

	return func(uri string) (*sql.DB, error) {
		return sql.Open(driver, uri)
	}
}
