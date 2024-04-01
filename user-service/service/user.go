package service

import (
	"context"
	pbu "user-service/genproto/user"
	l "user-service/pkg/logger"
	"user-service/storage"

	"github.com/jmoiron/sqlx"
	// "go.mongodb.org/mongo-driver/mongo"
)

// UserService ...
type UserService struct {
	storage storage.IStorage
	logger  l.Logger
}

// NewUserService ... Postgres
func NewUserService(db *sqlx.DB, log l.Logger) *UserService {
	return &UserService{
		storage: storage.NewStoragePg(db),
		logger:  log,
	}
}

// NewUserService ... Mongo
// func NewUserService(db *mongo.Client, log l.Logger) *UserService {
// 	return &UserService{
// 		storage: storage.NewStoragePg(db),
// 		logger:  log,
// 	}
// }

func (s *UserService) Login(ctx context.Context, req *pbu.LoginRequest) (*pbu.User, error) {
	user, err := s.storage.User().LoginDB(req)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) CheckUnique(ctx context.Context, req *pbu.CheckUniqueRequest) (*pbu.CheckUniqueRespons, error) {
	check, err := s.storage.User().CheckUniqueDB(req)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return check, nil
}

func (s *UserService) GetFullName(ctx context.Context, req *pbu.LoginRequest) (*pbu.User, error) {
	user, err := s.storage.User().GetFullNameDB(req)
	if err != nil {
		return nil, err
	}
	return user, nil
}


// User Start
func (s *UserService) CreateUser(ctx context.Context, req *pbu.User) (*pbu.User, error) {
	res, err := s.storage.User().CreateUserDB(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *UserService) ReadUser(ctx context.Context, req *pbu.IdRequest) (*pbu.User, error) {
	res, err := s.storage.User().ReadUserDB(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *UserService) UpdateUser(ctx context.Context, req *pbu.User) (*pbu.User, error) {
	res, err := s.storage.User().UpdateUserDB(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *UserService) DeleteUser(ctx context.Context, req *pbu.IdRequest) (*pbu.MessageResponse, error) {
	res, err := s.storage.User().DeleteUserDB(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *UserService) ListUser(ctx context.Context, req *pbu.GetAllRequest) (*pbu.ListUserResponse, error) {
	users, err := s.storage.User().ListUserDB(req)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *UserService) UpdatePassword(ctx context.Context, req *pbu.UpdatePasswordRequest) (*pbu.MessageResponse, error) {
	users, err := s.storage.User().UpdatePasswordDB(req)
	if err != nil {
		return nil, err
	}
	return users, nil
}
