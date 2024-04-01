package mocktest

import (
	pbu "api-gateway/genproto/user"
	"context"

	"google.golang.org/grpc"
)

type UserServiceClient interface {
	MockCreateUser(ctx context.Context, in *pbu.User, opts ...grpc.CallOption) (*pbu.User, error)
	MockReadUser(ctx context.Context, in *pbu.IdRequest, opts ...grpc.CallOption) (*pbu.User, error)
	MockUpdateUser(ctx context.Context, in *pbu.User, opts ...grpc.CallOption) (*pbu.User, error)
	MockDeleteUser(ctx context.Context, in *pbu.IdRequest, opts ...grpc.CallOption) (*pbu.MessageResponse, error)
	MockListUser(ctx context.Context, in *pbu.GetAllRequest, opts ...grpc.CallOption) (*pbu.ListUserResponse, error)
	// CheckField(ctx context.Context, in *pbu.CheckFieldRequest, opts ...grpc.CallOption) (*pbu.CheckFieldResponse, error)
}

type userServiceClient struct {
}

func (h *handlerV1) MockCreateUser(ctx context.Context, in *pbu.User, opts ...grpc.CallOption) (*pbu.User, error) {
	return in, nil
}

func (h *handlerV1) MockUpdateUser(ctx context.Context, in *pbu.User, opts ...grpc.CallOption) (*pbu.User, error) {
	return in, nil
}

func (h *handlerV1) MockReadUser(ctx context.Context, in *pbu.IdRequest, opts ...grpc.CallOption) (*pbu.User, error) {
	user := pbu.User{
		Id:       "e7cd4295-b99c-4aa0-bf10-7145ea81f472",
		FirstName:     "Oybek",
		LastName:     "Atamatov",
		Username: "oybekmdp",
		PhoneNumber: "+998999790445",
		Bio: "Never give up!",
		BirthDay: "2003-08-01",
		Email:    "oybekatamatov999@gmail.com",
		Avatar: "https://optimalw.com/wp-content/uploads/2015/10/sample-avatar-300x300.jpg",
		Password: "hello1234",
		RefreshToken: "mocktoken",
	}
	return &user, nil
}

func (h *handlerV1) MockListUser(ctx context.Context, in *pbu.GetAllRequest, opts ...grpc.CallOption) (*pbu.ListUserResponse, error) {
	user := pbu.User{
		Id:       "e7cd4295-b99c-4aa0-bf10-7145ea81f472",
		FirstName:     "Oybek",
		LastName:     "Atamatov",
		Username: "oybekmdp",
		PhoneNumber: "+998999790445",
		Bio: "Never give up!",
		BirthDay: "2003-08-01",
		Email:    "oybekatamatov999@gmail.com",
		Avatar: "https://optimalw.com/wp-content/uploads/2015/10/sample-avatar-300x300.jpg",
		Password: "hello1234",
		RefreshToken: "mocktoken",
	}
	resp := pbu.ListUserResponse{
		Users: []*pbu.User{
			&user,
			&user,
			&user,
			&user,
			&user,
			&user,
			&user,
		},
	}
	return &resp, nil
}

func (h *handlerV1) MockDeleteUser(ctx context.Context, in *pbu.IdRequest, opts ...grpc.CallOption) (*pbu.MessageResponse, error) {
	return &pbu.MessageResponse{Message: "User Succesfuly deleted!"}, nil
}

// func (h *handlerV1) CheckField(ctx context.Context, in *pbu.CheckFieldRequest, opts ...grpc.CallOption) (*pbu.CheckFieldResponse, error) {
// 	return &pbu.CheckFieldResponse{
// 		Status: false,
// 	}, nil
// }
