package service

import (
	pbc "comment-service/genproto/comment"
	l "comment-service/pkg/logger"
	"comment-service/storage"
	"context"

	"github.com/jmoiron/sqlx"
	// "go.mongodb.org/mongo-driver/mongo"
)

// CommentService ...
type CommentService struct {
	storage storage.IStorage
	logger  l.Logger
}

// NewUserService ...

// Postgres
func NewCommentService(db *sqlx.DB, log l.Logger) *CommentService {
	return &CommentService{
		storage: storage.NewStoragePg(db),
		logger:  log,
	}
}

// Mongo
// func NewCommentService(db *mongo.Client, log l.Logger) *CommentService {
// 	return &CommentService{
// 		storage: storage.NewStoragePg(db),
// 		logger:  log,
// 	}
// }

// Comment Start
func (s *CommentService) CreateComment(ctx context.Context, req *pbc.Comment) (*pbc.Comment, error) {
	res, err := s.storage.Comment().CreateCommentDB(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *CommentService) ReadComment(ctx context.Context, req *pbc.IdRequest) (*pbc.Comment, error) {
	res, err := s.storage.Comment().ReadCommentDB(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *CommentService) UpdateComment(ctx context.Context, req *pbc.Comment) (*pbc.Comment, error) {
	res, err := s.storage.Comment().UpdateCommentDB(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *CommentService) DeleteComment(ctx context.Context, req *pbc.IdRequest) (*pbc.MessageResponse, error) {
	res, err := s.storage.Comment().DeleteCommentDB(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *CommentService) ListComments(ctx context.Context, req *pbc.GetAllRequest) (*pbc.ListCommentResponse, error) {
	comments, err := s.storage.Comment().ListCommentsDB(req)
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func (s *CommentService) ListCommentsByProductId(ctx context.Context, req *pbc.IdRequest) (*pbc.ListCommentResponse, error) {
	comments, err := s.storage.Comment().ListCommentsByProductIdDB(req)
	if err != nil {
		return nil, err
	}
	return comments, nil
}

// Comment End