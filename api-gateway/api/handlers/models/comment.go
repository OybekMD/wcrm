package models

type Comment struct {
	Id        int64  `json:"id"`
	Content   string `json:"content"`
	UserId    string `json:"user_id"`
	ProductId int64  `json:"product_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type CommentList struct {
	Comments []*Comment `json:"comments"`
}

type CommentReq struct {
	Content   string `json:"content"`
	UserId    string `json:"user_id"`
	ProductId int64  `json:"product_id"`
}

type CommentsByProductIdReq struct {
	Id    string `json:"id"`
	Page  int64  `json:"page"`
	Limit int64  `json:"limit"`
}

type CommentUpdate struct {
	Id      int64  `json:"id"`
	Content string `json:"content"`
}
