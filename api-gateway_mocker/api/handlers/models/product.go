package models

type Product struct {
	Id          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	DeletedAt   string `json:"deleted_at"`
}
type ProductReq struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
}

type ProductUpdate struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
}
