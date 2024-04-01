package models

type Orderproduct struct {
	Id        int64  `json:"id"`
	UserId    string `json:"user_id"`
	PictureId int64  `json:"picture_id"`
	CreatedAt string `json:"created_at"`
}
type OrderproductReq struct {
	UserId    string `json:"user_id"`
	PictureId int64  `json:"picture_id"`
}

type OrderproductUpdate struct {
	UserId    string `json:"user_id"`
	PictureId int64  `json:"picture_id"`
}
