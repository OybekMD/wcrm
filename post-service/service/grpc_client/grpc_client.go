package grpcClient

import (
	"fmt"
	"post-service/config"
	pbp "post-service/genproto/comment"

	"google.golang.org/grpc"
)

type IServiceManager interface {
	CommentService() pbp.CommentServiceClient
}

type serviceManager struct {
	cfg            config.Config
	commentService    pbp.CommentServiceClient
}

func New(cfg config.Config) (IServiceManager, error) {
	// dail to course-service
	connComment, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.CommentServiceHost, cfg.CommentServicePort),
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, fmt.Errorf("user service dail host: %s port : %d", cfg.CommentServiceHost, cfg.CommentServicePort)
	}
	return &serviceManager{
		cfg:            cfg,
		commentService:    pbp.NewCommentServiceClient(connComment),
	}, nil
}

func (s *serviceManager) CommentService() pbp.CommentServiceClient {
	return s.commentService
}
