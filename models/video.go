package models

import (
	"database/sql"
	"errors"
)

type Video struct {
	BaseModel
	Title     string
	UserID    int
	Cover     string
	URL       string
	Likes     uint
	Dislikes  uint
	CreatedAt string
}

type VideoStorage struct {
	Storage
}

func NewVideoStorage(conn *sql.DB) *VideoStorage {

	return &VideoStorage{
		Storage: Storage{
			conn: conn,
		},
	}
}

func (storage *VideoStorage) GetConnection(conn *sql.DB) {
	storage.conn = conn
}

func (storage *VideoStorage) Find(pointer interface{}) error {

	if videos, ok := pointer.(*[]Video); ok {
		stmt, err := storage.conn.Prepare("SELECT * FROM videos WHERE   ")
		CheckError(err)

		rows, err := stmt.Query()
		CheckError(err)
		for rows.Next() {
			video := Video{}
			err := rows.Scan(&video.ID, &video.URL, &video.Likes, &video.Dislikes, &video.UserID, &video.Title, &video.Cover, &video.CreatedAt)
			CheckError(err)

			*videos = append(*videos, video)
		}

		pointer = &videos
		return nil

	}

	return errors.New("Invalid user slice pointer ")

}

func (storage *VideoStorage) FindOne(condition map[string]string, pointer interface{}) error {

	stmt, err := storage.conn.Prepare("SELECT * FROM videos WHERE id = ")
	CheckError(err)
	rows, err := stmt.Query(condition["id"])
	CheckError(err)

	if video, ok := pointer.(*Video); ok {
		for rows.Next() {
			err := rows.Scan(&video.ID, &video.URL, &video.Likes, &video.Dislikes, &video.UserID, &video.Title, &video.Cover, &video.CreatedAt)
			CheckError(err)

			pointer = video
			break
		}

	}

	return nil
}

func (storage *VideoStorage) Save(data interface{}) error {
	if video, ok := data.(*Video); ok {
		stmt, err := storage.conn.Prepare("INSERT INTO videos (url, likes, dislikes, user_id,  title, cover) VALUES (? ,? , ? , ?, ? , ? )")
		CheckError(err)

		result, err := stmt.Exec(video.URL, &video.Likes, &video.Dislikes, &video.UserID, &video.Title, &video.Cover)
		CheckError(err)

		_, err = result.LastInsertId()
		CheckError(err)

	}
	return nil

}

func (storage *VideoStorage) Update(condition map[string]string, data interface{}) error {

	if video, ok := data.(Video); ok {
		stmt, err := storage.conn.Prepare("UPDATE videos SET url = ?, likes = ?, dislikes = ?, user_id = ?,  title = ?, cover = ?   WHERE id =  ? ")
		CheckError(err)
		result, err := stmt.Exec(&video.URL, &video.Likes, &video.Dislikes, &video.UserID, &video.Title, &video.Cover, condition["id"])
		CheckError(err)

		_, err = result.LastInsertId()
		CheckError(err)

	}

	return nil
}

func (storage *VideoStorage) Delete(condition map[string]string) bool {
	stmt, err := storage.conn.Prepare("DELETE FROM videos WHERE id = ?")
	CheckError(err)
	result, err := stmt.Exec(condition["id"])
	count, err := result.RowsAffected()
	CheckError(err)

	return count > 0
}
