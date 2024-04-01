package models

// Error ...
type Error struct {
	Message string `json:"message"`
}

// StandardErrorModel ...
type StandardErrorModel struct {
	Error Error `json:"error"`
}

type ResponseError struct {
	Status  error  `json:"status"`
	Message string `json:"message"`
}
