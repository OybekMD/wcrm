package postgres

import (
	"github.com/jmoiron/sqlx"
)

type postRepo struct {
	db *sqlx.DB
}

// NewUserRepo ...
func NewPostRepo(db *sqlx.DB) *postRepo {
	return &postRepo{db: db}
}
