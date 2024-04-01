package postgres

import (
	pbu "user-service/genproto/user"

	"github.com/jmoiron/sqlx"
)

type userRepoMock struct {
	db *sqlx.DB
}

// NewUserRepo ...
func NewUserRepoMock(db *sqlx.DB) *userRepoMock {
	return &userRepoMock{db: db}
}

func (r *userRepoMock) CreateUserDB(req *pbu.User) (*pbu.User, error) {

	return &pbu.User{
		Id:           req.Id,
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		Username:     req.Username,
		PhoneNumber:  req.PhoneNumber,
		Bio:          req.Bio,
		BirthDay:     req.BirthDay,
		Email:        req.Email,
		Avatar:       req.Avatar,
		Password:     req.Password,
		RefreshToken: "testToken",
		CreatedAt:    "2003-08-01",
		UpdatedAt:    "2003-08-01",
	}, nil
}

func (r *userRepoMock) ReadUserDB(req *pbu.IdRequest) (*pbu.User, error) {

	return &pbu.User{
		Id:           req.Id,
		FirstName:    "Oybek",
		LastName:     "Atamatov",
		Username:     "oybekmd",
		PhoneNumber:  "998999790445",
		Bio:          "Test bio",
		BirthDay:     "2003-08-01",
		Email:        "testemail@gamil.com",
		Avatar:       "example.com/user.jpg",
		Password:     "hello1234",
		RefreshToken: "testToken",
	}, nil
}

func (r *userRepoMock) UpdateUserDB(req *pbu.User) (*pbu.User, error) {

	return &pbu.User{
		Id:           req.Id,
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		Username:     req.Username,
		PhoneNumber:  req.PhoneNumber,
		Bio:          req.Bio,
		BirthDay:     req.BirthDay,
		Email:        req.Email,
		Avatar:       req.Avatar,
		Password:     req.Password,
		RefreshToken: "testToken",
		CreatedAt:    "2003-08-01",
		UpdatedAt:    "2003-08-01",
	}, nil

}

func (r *userRepoMock) DeleteUserDB(req *pbu.IdRequest) error {
	return nil

}

func (r *userRepoMock) ListUserDB(req *pbu.GetAllRequest) (*pbu.ListUserResponse, error) {
	var allUser pbu.ListUserResponse
	u1 := pbu.User{
		Id:           "0e7de31f-7994-404a-a453-ecceba69ffdd",
		FirstName:    "Xasan",
		LastName:     "Nosirov",
		Username:     "xasannosirov",
		PhoneNumber:  "998990010011",
		Bio:          "Learn and Pushlish",
		BirthDay:     "2006-01-01",
		Email:        "xasannosirov@gmail.com",
		Avatar:       "example.com/user.jpg",
		Password:     "StrongPassword",
		RefreshToken: "testToken",
		CreatedAt:    "2023-01-01",
		UpdatedAt:    "2024-01-01",
	}
	u2 := pbu.User{
		Id:           "0e7de31f-7994-404a-a453-ecceba69ffdd",
		FirstName:    "Nodirbek",
		LastName:     "Nomonov",
		Username:     "rarebek",
		PhoneNumber:  "998990010011",
		Bio:          "Learn and Pushlish",
		BirthDay:     "2006-01-01",
		Email:        "rare@gmail.com",
		Avatar:       "example.com/user.jpg",
		Password:     "StrongPassword",
		RefreshToken: "testToken",
		CreatedAt:    "2023-01-01",
		UpdatedAt:    "2024-01-01",
	}
	allUser.Users = append(allUser.Users, &u1)
	allUser.Users = append(allUser.Users, &u2)

	return &allUser, nil
}
