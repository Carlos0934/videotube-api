package models

import (
	"database/sql"
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

type ICRUDModel interface {
	Find(pointer interface{}) error
	FindOne(codition map[string]string, pointer interface{}) error
	Save(data interface{}) error
	Update(condition map[string]string, data interface{}) error
	Delete(condition map[string]string) bool
}
