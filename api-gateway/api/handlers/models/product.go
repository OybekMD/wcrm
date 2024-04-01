package models

type Product struct {
	Id          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
	Picture     string `json:"picture"`
	CategoryId  int64  `json:"category_id"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	DeletedAt   string `json:"deleted_at"`
}

type ProductReq struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
	Picture     string `json:"picture"`
	CategoryId  int64  `json:"category_id"`
}

type ProductUpdate struct {
	Id          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
	CategoryId  int64  `json:"category_id"`
}
