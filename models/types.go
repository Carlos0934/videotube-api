package models

import "database/sql"

type IStorage interface {
	GetConnection(*sql.DB)
}

type ICRUDModel interface {
	Find(condition map[string]string, pointer interface{}) error
	FindOne(codition map[string]string, pointer interface{}) error
	Save(data interface{}) error
	Update(condition map[string]string, data interface{}) error
	Delete(condition map[string]string) bool
}
