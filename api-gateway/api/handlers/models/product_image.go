package models

type ProductImage struct {
	Id        int64  `json:"id"`
	Picture   string `json:"picture"`
	ProductId int64  `json:"product_id"`
	CreatedAt string `json:"created_at"`
	DeletedAt string `json:"deleted_at"`
}
type ProductImageReq struct {
	Picture   string `json:"picture"`
	ProductId int64  `json:"product_id"`
}

type ProductImageUpdate struct {
	Picture   string `json:"picture"`
	ProductId int64  `json:"product_id"`
}
