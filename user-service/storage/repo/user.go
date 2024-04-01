package repo

import (
	pbu "user-service/genproto/user"
)

// UserStorageI ...
type UserStorageI interface {
	LoginDB(*pbu.LoginRequest) (*pbu.User, error)
	CheckUniqueDB(req *pbu.CheckUniqueRequest) (*pbu.CheckUniqueRespons, error)
	GetFullNameDB(req *pbu.LoginRequest) (*pbu.User, error)

	CreateUserDB(*pbu.User) (*pbu.User, error)
	ReadUserDB(*pbu.IdRequest) (*pbu.User, error)
	UpdateUserDB(*pbu.User) (*pbu.User, error)
	DeleteUserDB(*pbu.IdRequest) (*pbu.MessageResponse, error)
	ListUserDB(*pbu.GetAllRequest) (*pbu.ListUserResponse, error)
	UpdatePasswordDB(*pbu.UpdatePasswordRequest) (*pbu.MessageResponse, error)
}
