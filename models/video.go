package models

type Video struct {
	BaseModel
	Title     string
	UserID    string
	Cover     string
	URL       string
	Likes     uint
	Dislikes  uint
	CreatedAt string
}

type VideoStorage struct {
}
