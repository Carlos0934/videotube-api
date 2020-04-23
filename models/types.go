package models

import (
	"database/sql"
	"fmt"
	"strings"
)

type IStorage interface {
	GetConnection(*sql.DB)
}
type BaseModel struct {
	ID int
}

type Storage struct {
	conn *sql.DB
}

func (storage *Storage) ParseQuery(query string, condition map[string]interface{}) (string, []interface{}) {
	parseQuery := query + " WHERE"
	params := make([]interface{}, 0)

	for key, value := range condition {
		parseQuery += fmt.Sprintf(" %v = ? AND ", key)
		params = append(params, value)
	}
	index := strings.LastIndex(parseQuery, "AND")

	parseQuery = parseQuery[:index]

	return parseQuery, params
}

type ICRUDModel interface {
	Find(pointer interface{}) error
	FindOne(codition map[string]interface{}, pointer interface{}) error
	Save(data interface{}) error
	Update(condition map[string]interface{}, data interface{}) error
	Delete(condition map[string]interface{}) bool
}
