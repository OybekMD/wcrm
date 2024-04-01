package models

type Post struct {
	Id        int64  `json:"id"`
	Picture   string `json:"picture"`
	Title     string `json:"title"`
	Article   string `json:"article"`
	OwnerId   string `json:"owner_id"`
	CreatedAt string `json:"created_at"`
	UpdetedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}
