package models

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	BaseModel
	Username  string
	Email     string
	Password  string
	Birthdate string
}

type UserStorage struct {
	Storage
}

func NewUserStorage(conn *sql.DB) *UserStorage {

	return &UserStorage{
		Storage: Storage{
			conn: conn,
		},
	}

}

func (storage *UserStorage) GetConnection(conn *sql.DB) {
	storage.conn = conn
}

func (storage *UserStorage) Find(pointer interface{}) error {
	stmt, err := storage.conn.Prepare("SELECT * FROM users  ")
	CheckError(err)

	rows, err := stmt.Query()
	CheckError(err)

	if users, ok := pointer.(*[]User); ok {
		for rows.Next() {
			user := User{}
			err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Birthdate)
			CheckError(err)

			*users = append(*users, user)
		}

		pointer = &users
		return nil

	}

	return errors.New("Invalid user slice pointer ")

}

func (storage *UserStorage) FindOne(condition map[string]string, pointer interface{}) error {

	stmt, err := storage.conn.Prepare("SELECT * FROM users WHERE username = ? AND password = ?")
	CheckError(err)
	rows, err := stmt.Query(condition["username"], condition["password"])
	CheckError(err)

	if user, ok := pointer.(*User); ok {
		for rows.Next() {
			err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Birthdate)
			CheckError(err)

			pointer = user
			break
		}

	}

	return nil
}

func (storage *UserStorage) Save(data interface{}) error {
	if user, ok := data.(*User); ok {
		stmt, err := storage.conn.Prepare("INSERT INTO users (username, email, password, birthdate) VALUES (? ,? , ? , ? )")
		CheckError(err)

		result, err := stmt.Exec(user.Username, user.Email, user.Password, user.Birthdate)
		CheckError(err)

		_, err = result.LastInsertId()
		CheckError(err)

		return err
		fmt.Println("dsadsa")
	}
	return nil

}

func (storage *UserStorage) Update(condition map[string]string, data interface{}) error {

	if user, ok := data.(User); ok {
		stmt, err := storage.conn.Prepare("UPDATE users SET  username = ? , email = ? , password = ? ,  birthdate = ? WHERE id =  ? ")
		CheckError(err)
		result, err := stmt.Exec(user.Username, user.Email, user.Password, user.Birthdate, condition["id"])
		CheckError(err)

		_, err = result.LastInsertId()
		CheckError(err)

	}

	return nil
}

func (storage *UserStorage) Delete(condition map[string]string) bool {
	stmt, err := storage.conn.Prepare("DELETE FROM users WHERE id = ?")
	CheckError(err)
	result, err := stmt.Exec(condition["id"])
	count, err := result.RowsAffected()
	CheckError(err)

	return count > 0
}
