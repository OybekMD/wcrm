package postgres

import (
	pbc "comment-service/genproto/comment"

	"github.com/jmoiron/sqlx"
)

type commentRepo struct {
	db *sqlx.DB
}

// NewCommentRepo ...
func NewCommentRepo(db *sqlx.DB) *commentRepo {
	return &commentRepo{db: db}
}

// -------------------------------------------
// Comment CRUD with Soft Delete
// -------------------------------------------

func (r *commentRepo) CreateCommentDB(req *pbc.Comment) (*pbc.Comment, error) {
	var res pbc.Comment
	query := `
        INSERT INTO comments(
            content,
            user_id,
            product_id
        )
        VALUES ($1, $2, $3) 
        RETURNING 
            id,
            content,
            user_id,
            product_id,
            created_at,
            updated_at
    `
	err := r.db.QueryRow(
		query,
		req.Content,
		req.UserId,
		req.ProductId).Scan(
		&res.Id,
		&res.Content,
		&res.UserId,
		&res.ProductId,
		&res.CreatedAt,
		&res.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (r *commentRepo) ReadCommentDB(req *pbc.IdRequest) (*pbc.Comment, error) {
	var res pbc.Comment
	query := `
        SELECT
			id,
			content,
			user_id,
			product_id,
			created_at,
			updated_at
        FROM 
            comments
        WHERE
            id=$1 AND
            deleted_at IS NULL
    `
	err := r.db.QueryRow(query, req.Id).Scan(
		&res.Id,
		&res.Content,
		&res.UserId,
		&res.ProductId,
		&res.CreatedAt,
		&res.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (r *commentRepo) UpdateCommentDB(req *pbc.Comment) (*pbc.Comment, error) {
	var res pbc.Comment
	query := `
        UPDATE
            comments
        SET
			content=$1,
            updated_at=CURRENT_TIMESTAMP
        WHERE
            id=$2
        RETURNING
			id,
			content,
			user_id,
			product_id,
			created_at,
			updated_at
    `
	err := r.db.QueryRow(
		query,
		req.Content,
		req.Id).Scan(
		&res.Id,
		&res.Content,
		&res.UserId,
		&res.ProductId,
		&res.CreatedAt,
		&res.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (r *commentRepo) DeleteCommentDB(req *pbc.IdRequest) (*pbc.MessageResponse, error) {
	var res pbc.MessageResponse

	query := `
        UPDATE
            comments
        SET
            deleted_at=CURRENT_TIMESTAMP
        WHERE
            id=$1
    `

	result, err := r.db.Exec(query, req.Id)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		// No rows were affected, indicating that the record was not found
		res.Message = "Comment not found for deletion."
	} else {
		res.Message = "Comment successfully deleted!"
	}

	return &res, nil
}

func (r *commentRepo) ListCommentsDB(req *pbc.GetAllRequest) (*pbc.ListCommentResponse, error) {
	var allComments pbc.ListCommentResponse
	query := `
        SELECT
			id,
			content,
			user_id,
			product_id,
			created_at,
			updated_at
        FROM 
            comments
        WHERE
            deleted_at IS NULL
        LIMIT $1
        OFFSET $2
    `
	offset := req.Limit * (req.Page - 1)
	rows, err := r.db.Query(query, req.Limit, offset)
	if err != nil {
		return nil, err
	}
	
	for rows.Next() {
		var comment pbc.Comment
		err := rows.Scan(
			&comment.Id,
			&comment.Content,
			&comment.UserId,
			&comment.ProductId,
			&comment.CreatedAt,
			&comment.UpdatedAt)
		if err != nil {
			return nil, err
		}
		allComments.Comments = append(allComments.Comments, &comment)
	}
	return &allComments, nil
}

func (r *commentRepo) ListCommentsByProductIdDB(req *pbc.ListPorductIdRequest) (*pbc.ListCommentResponse, error) {
	var allComments pbc.ListCommentResponse
	query := `
        SELECT
			id,
			content,
			user_id,
			product_id,
			created_at,
			updated_at
        FROM 
            comments
        WHERE
			product_id = $1
            deleted_at IS NULL
        LIMIT $2
        OFFSET $3
    `
	offset := req.Limit * (req.Page - 1)
	rows, err := r.db.Query(query, req.Id, req.Limit, offset)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var comment pbc.Comment
		err := rows.Scan(
			&comment.Id,
			&comment.Content,
			&comment.UserId,
			&comment.ProductId,
			&comment.CreatedAt,
			&comment.UpdatedAt)
		if err != nil {
			return nil, err
		}
		allComments.Comments = append(allComments.Comments, &comment)
	}
	return &allComments, nil
}
