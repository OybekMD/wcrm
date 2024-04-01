package models

type Comment struct {
	Id        int64  `json:"id"`
	Content   string `json:"content"`
	UserId    string `json:"user_id"`
	ProductId int64  `json:"product_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}

type CommentReq struct {
	Content   string `json:"content"`
	UserId    string `json:"user_id"`
	ProductId int64  `json:"product_id"`
}

type CommentUpdate struct {
	Id        int64  `json:"id"`
	Content   string `json:"content"`
}
