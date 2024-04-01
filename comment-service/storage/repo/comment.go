package repo

import (
	pbc "comment-service/genproto/comment"
)

// CommentStorageI ...
type CommentStorageI interface {
	CreateCommentDB(*pbc.Comment) (*pbc.Comment, error)
	ReadCommentDB(*pbc.IdRequest) (*pbc.Comment, error)
	UpdateCommentDB(*pbc.Comment) (*pbc.Comment, error)
	DeleteCommentDB(*pbc.IdRequest) (*pbc.MessageResponse, error)
	ListCommentsDB(*pbc.GetAllRequest) (*pbc.ListCommentResponse, error)

	ListCommentsByProductIdDB(*pbc.ListPorductIdRequest) (*pbc.ListCommentResponse, error)
}
