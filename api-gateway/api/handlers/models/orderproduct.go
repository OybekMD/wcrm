package models

type Orderproduct struct {
	Id        int64  `json:"id"`
	UserId    string `json:"user_id"`
	ProductId int64  `json:"product_id"`
	CreatedAt string `json:"created_at"`
}
type OrderproductReq struct {
	UserId    string `json:"user_id"`
	ProductId int64  `json:"product_id"`
}

type OrderproductUpdate struct {
	UserId    string `json:"user_id"`
	ProductId int64  `json:"product_id"`
}
