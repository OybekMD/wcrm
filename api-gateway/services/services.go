package services

import (
	"fmt"

	// mocktest "api-gateway/api_mock/mock-test"
	"api-gateway/config"
	pbc "api-gateway/genproto/comment"
	pbp "api-gateway/genproto/post"
	pbu "api-gateway/genproto/user"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
)

type IServiceManager interface {
	UserService() pbu.UserServiceClient
	PostService() pbp.PostServiceClient
	CommentService() pbc.CommentServiceClient
	// MockUserService() mocktest.UserServiceClient
}

type serviceManager struct {
	userService pbu.UserServiceClient
	// mockUserService mocktest.UserServiceClient
	postService    pbp.PostServiceClient
	commentService pbc.CommentServiceClient
}

func (s *serviceManager) UserService() pbu.UserServiceClient {
	return s.userService
}

// func (s *serviceManager) MockUserService() mocktest.UserServiceClient {
// 	return s.mockUserService
// }

func (s *serviceManager) PostService() pbp.PostServiceClient {
	return s.postService
}

func (s *serviceManager) CommentService() pbc.CommentServiceClient {
	return s.commentService
}

func NewServiceManager(conf *config.Config) (IServiceManager, error) {
	resolver.SetDefaultScheme("dns")

	connUser, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.UserServiceHost, conf.UserServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	connPost, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.PostServiceHost, conf.PostServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	connComment, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.CommentServiceHost, conf.CommentServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	serviceManager := &serviceManager{
		userService: pbu.NewUserServiceClient(connUser),
		// mockUserService: mocktest.NewUserServiceClient(),
		postService: pbp.NewPostServiceClient(connPost),
		commentService: pbc.NewCommentServiceClient(connComment),
	}

	return serviceManager, nil
}
