package models

type Category struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	UserId    string `json:"user_id"`
	IconId    int64  `json:"icon_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}
type CategoryReq struct {
	Name   string `json:"name"`
	UserId string `json:"user_id"`
	IconId int64  `json:"icon_id"`
}

type CategoryUpdate struct {
	Name   string `json:"name"`
	UserId string `json:"user_id"`
	IconId int64  `json:"icon_id"`
}
