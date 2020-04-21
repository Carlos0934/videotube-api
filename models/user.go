package models

import "database/sql"

type User struct {
	BaseModel
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
	stmt, err := storage.conn.Prepare("SELECT * FROM USERS WHERE ")
	CheckError(err)

	rows, err := stmt.Query()
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
	stmt, err := storage.conn.Prepare("SELECT * FROM USERS WHERE username = ? AND password = ?")

	rows, err := stmt.Query(condition["username"], condition["password"])
	CheckError(err)
	var user *User
	rows.Scan(user.ID, user.Username, user.Email, user.Password, user.Birthdate)
	pointer = user
	return nil
}

func (storage *UserStorage) Save(data interface{}) error {
	if user, ok := data.(User); ok {
		stmt, err := storage.conn.Prepare("INSERT INTO USERS (username, email, password, birthdate) VALUES (? ,? , ? , ? )")
		CheckError(err)
		result, err := stmt.Exec(user.Username, user.Email, user.Password, user.Birthdate)
		CheckError(err)

		_, err = result.LastInsertId()
		CheckError(err)

	}
	return nil

}

func (storage *UserStorage) Update(condition map[string]string, data interface{}) error {

}

func (storage *UserStorage) Delete(condition map[string]string) bool {

}
