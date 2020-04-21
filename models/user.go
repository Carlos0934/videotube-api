package models

import "database/sql"

type User struct {
	ID        string
	Username  string
	Email     string
	Password  string
	Birthdate string
}

type UserStorage struct {
	conn *sql.DB
}

func (storage *UserStorage) GetConnection(conn *sql.DB) {
	storage.conn = conn
}

func (storage *UserStorage) Find(condition map[string]string, pointer interface{}) error {
	stmt, err := storage.conn.Prepare("SELECT * FROM USERS WHERE username = ? AND password = ?")
	CheckError(err)

	rows, err := stmt.Query(condition["username"], condition["password"])
	CheckError(err)
	users := make([]*User, 0)
	for rows.Next() {
		var user *User
		err := rows.Scan(user.ID, user.Username, user.Email, user.Password, user.Birthdate)
		CheckError(err)
		users = append(users, user)
	}
	pointer = users
	return nil
}

func (storage *UserStorage) FindOne(condition map[string]string, pointer interface{}) error {

}

func (storage *UserStorage) Save(data interface{}) error {

}

func (storage *UserStorage) Update(condition map[string]string, data interface{}) error {

}

func (storage *UserStorage) Delete(condition map[string]string) bool {

}
