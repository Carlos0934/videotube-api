package models

type Video struct {
	BaseModel
	Title    string
	UserID   string
	Cover    string
	Url      string
	Likes    uint
	Dislikes uint
}

type VideoStorage struct {
}
