package models

import (
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

type Video struct {
	BaseModel
	Title     string  `json:"title"`
	UserID    int     `json:"user_id"`
	Cover     *string `json:"cover"`
	URL       string  `json:"url"`
	Likes     uint    `json:"likes"`
	Dislikes  uint    `json:"dislikes"`
	CreatedAt string  `json:"created_at"`
}

const (
	ContentCover string = "cover"
	ContentVideo string = "video"
)

type VideoStorage struct {
	Storage
	CoverDir string
	VideoDir string
}

func NewVideoStorage(conn *sql.DB) *VideoStorage {

	return &VideoStorage{
		Storage: Storage{
			conn: conn,
		},

		CoverDir: "./data/covers/",
		VideoDir: "./data/videos/",
	}
}

func (storage *VideoStorage) GetConnection(conn *sql.DB) {
	storage.conn = conn
}

func (storage *VideoStorage) Find(pointer interface{}) error {

	if videos, ok := pointer.(*[]Video); ok {

		stmt, err := storage.conn.Prepare("SELECT * FROM videos  ")
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
		return err

	}

	return videoErr

}

func (storage *VideoStorage) FindOne(condition map[string]interface{}, pointer interface{}) error {

	query, data := storage.ParseQuery("SELECT * FROM videos", condition)
	stmt, err := storage.conn.Prepare(query)
	CheckError(err)
	rows, err := stmt.Query(data...)
	CheckError(err)

	if video, ok := pointer.(*Video); ok {
		for rows.Next() {
			err := rows.Scan(&video.ID, &video.URL, &video.Likes, &video.Dislikes, &video.UserID, &video.Title, &video.Cover, &video.CreatedAt)
			CheckError(err)

			pointer = video
			break
		}

		return err

	}

	return videoErr
}

func (storage *VideoStorage) FindByUser(id string, pointer interface{}) error {
	if videos, ok := pointer.(*[]Video); ok {
		stmt, err := storage.conn.Prepare("SELECT * FROM videos WHERE user_id = ?  ")
		CheckError(err)

		rows, err := stmt.Query(id)
		CheckError(err)
		for rows.Next() {
			video := Video{}
			err := rows.Scan(&video.ID, &video.URL, &video.Likes, &video.Dislikes, &video.UserID, &video.Title, &video.Cover, &video.CreatedAt)
			CheckError(err)

			*videos = append(*videos, video)

		}

		pointer = &videos

		return err
	}

	return videoErr
}
func (storage *VideoStorage) Save(data interface{}) error {
	if video, ok := data.(*Video); ok {

		stmt, err := storage.conn.Prepare("INSERT INTO videos ( user_id,  title) VALUES (? ,? )")
		CheckError(err)

		result, err := stmt.Exec(&video.UserID, &video.Title)
		CheckError(err)

		_, err = result.LastInsertId()
		CheckError(err)

		return err

	}
	return videoErr

}
func (storage *VideoStorage) getPath(filename string, context string) (string, error) {
	path := ""
	if context == "cover" {
		path = storage.CoverDir + filename
	} else if context == "video" {
		path = storage.VideoDir + filename
	} else {

		return "", errors.New("Invalid context")
	}

	return path, nil
}

func (Storage *VideoStorage) ParseFilename(filename string) string {
	hash := sha256.Sum256([]byte(filename + time.Now().String()))

	return base64.URLEncoding.EncodeToString(hash[:]) + filepath.Ext(filename)
}
func (storage *VideoStorage) SerializeContent(content []byte, filename string, context string) string {

	newFilename := storage.ParseFilename(filename)

	path, err := storage.getPath(newFilename, context)
	CheckError(err)

	file, err := os.Create(path)
	CheckError(err)
	_, err = file.Write(content)
	CheckError(err)
	err = file.Close()
	CheckError(err)
	return newFilename
}

func (storage *VideoStorage) DeserializeContent(filename string, context string) []byte {
	path, err := storage.getPath(filename, context)
	CheckError(err)

	file, err := ioutil.ReadFile(path)
	CheckError(err)
	return file
}

func (storage *VideoStorage) Update(condition map[string]interface{}, data interface{}) error {

	if video, ok := data.(*Video); ok {
		stmt, err := storage.conn.Prepare("UPDATE videos SET  likes = ?, dislikes = ?,   title = ?    WHERE id =  ? ")
		CheckError(err)
		result, err := stmt.Exec(&video.Likes, &video.Dislikes, &video.Title, condition["id"])
		CheckError(err)

		_, err = result.LastInsertId()
		CheckError(err)
		return err
	}

	return videoErr
}

func (storage *VideoStorage) AddURLS(id string, data [2]string) error {
	stmt, err := storage.conn.Prepare("UPDATE videos SET  url = ?,  cover = ?  WHERE id =  ?  ")
	CheckError(err)
	result, err := stmt.Exec(&data[0], &data[1], &id)
	CheckError(err)

	_, err = result.LastInsertId()
	CheckError(err)
	return err
}

func (storage *VideoStorage) deleteContent(condition map[string]interface{}) {
	stmt, err := storage.conn.Prepare("SELECT url, cover FROM videos WHERE id = ? LIMIT 1")
	CheckError(err)
	row, err := stmt.Query(condition["id"])
	var (
		cover string
		video string
	)
	for row.Next() {

		err := row.Scan(&video, &cover)
		CheckError(err)
	}
	video = storage.VideoDir + filepath.Base(video)
	cover = storage.CoverDir + filepath.Base(cover)

	err = os.Remove(cover)
	CheckError(err)
	err = os.Remove(video)
	CheckError(err)
}
func (storage *VideoStorage) Delete(condition map[string]interface{}) bool {
	storage.deleteContent(condition)
	stmt, err := storage.conn.Prepare("DELETE FROM videos WHERE id = ?")
	CheckError(err)
	result, err := stmt.Exec(condition["id"])
	count, err := result.RowsAffected()
	CheckError(err)

	return count > 0
}
