package postgres

import (
	pbp "post-service/genproto/post"
)

// -------------------------------------------
// Category CRUD with Soft Delete
// -------------------------------------------

func (r *postRepo) CreateCategoryDB(req *pbp.Category) (*pbp.Category, error) {
	var res pbp.Category
	query := `
        INSERT INTO categorys(
            name,
            icon_id
        )
        VALUES ($1, $2) 
        RETURNING 
            id,
            name,
            icon_id,
            created_at,
            updated_at
    `
	err := r.db.QueryRow(
		query,
		req.Name,
		req.IconId).Scan(
		&res.Id,
		&res.Name,
		&res.IconId,
		&res.CreatedAt,
		&res.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (r *postRepo) ReadCategoryDB(req *pbp.IdRequest) (*pbp.Category, error) {
	var res pbp.Category
	query := `
        SELECT
            id,
            name,
            icon_id,
            created_at,
            updated_at
        FROM 
            categorys
        WHERE
            id=$1 AND
            deleted_at IS NULL
    `
	err := r.db.QueryRow(query, req.Id).Scan(
		&res.Id,
		&res.Name,
		&res.IconId,
		&res.CreatedAt,
		&res.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (r *postRepo) UpdateCategoryDB(req *pbp.Category) (*pbp.Category, error) {
	var res pbp.Category
	query := `
        UPDATE
            categorys
        SET
            name=$1,
            icon_id=$2,
            updated_at=CURRENT_TIMESTAMP
        WHERE
            id=$3
        RETURNING
            id,
            name,
            icon_id,
            created_at,
            updated_at
    `
	err := r.db.QueryRow(
		query,
		req.Name,
		req.IconId,
		req.Id).Scan(
		&res.Id,
		&res.Name,
		&res.IconId,
		&res.CreatedAt,
		&res.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (r *postRepo) DeleteCategoryDB(req *pbp.IdRequest) (*pbp.MessageResponse, error) {
	var res pbp.MessageResponse

	query := `
        UPDATE
            categorys
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
		res.Message = "Category not found for deletion."
	} else {
		res.Message = "Category successfully deleted!"
	}

	return &res, nil
}

func (r *postRepo) ListCategorysDB(req *pbp.GetAllRequest) (*pbp.ListCategoryResponse, error) {
	var allCategories pbp.ListCategoryResponse
	query := `
        SELECT
            id, 
            name,
            icon_id,
            created_at,
            updated_at
        FROM 
            categorys
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
		var category pbp.Category
		err := rows.Scan(
			&category.Id,
			&category.Name,
			&category.IconId,
			&category.CreatedAt,
			&category.UpdatedAt)
		if err != nil {
			return nil, err
		}
		allCategories.Categorys = append(allCategories.Categorys, &category)
	}
	return &allCategories, nil
}
