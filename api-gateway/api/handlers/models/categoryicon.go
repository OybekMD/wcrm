package models

type CategoryIcon struct {
	Id      int64  `json:"id"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
}

type CategoryIconReq struct {
	Name    string `json:"name"`
	Picture string `json:"picture"`
}

type CategoryIconUpdate struct {
	Id      int64  `json:"id"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
}
