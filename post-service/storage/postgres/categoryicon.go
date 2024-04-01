package postgres

import (
	pbp "post-service/genproto/post"
)

// -------------------------------------------
// CategoryIcon CRUD
// -------------------------------------------

func (r *postRepo) CreateCategoryIconDB(req *pbp.CategoryIcon) (*pbp.CategoryIcon, error) {
	var res pbp.CategoryIcon
	query := `
		INSERT INTO category_icons(
			name,
			picture
		)
		VALUES ($1, $2) 
		RETURNING 
			id,
			name,
			picture
	`
	err := r.db.QueryRow(
		query,
		req.Name,
		req.Picture).Scan(
		&res.Id,
		&res.Name,
		&res.Picture)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (r *postRepo) ReadCategoryIconDB(req *pbp.IdRequest) (*pbp.CategoryIcon, error) {
	var res pbp.CategoryIcon
	query := `
	SELECT
		id,
		name,
		picture
	FROM 
		category_icons
	WHERE
		id=$1
	`
	err := r.db.QueryRow(query, req.Id).Scan(
		&res.Id,
		&res.Name,
		&res.Picture)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (r *postRepo) UpdateCategoryIconDB(req *pbp.CategoryIcon) (*pbp.CategoryIcon, error) {
	var res pbp.CategoryIcon
	query := `
	UPDATE
		category_icons
	SET
		name=$1,
		picture=$2
	WHERE
		id=$3
	RETURNING
		id,
		name,
		picture
	`
	err := r.db.QueryRow(
		query,
		req.Name,
		req.Picture,
		req.Id).Scan(
		&res.Id,
		&res.Name,
		&res.Picture)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (r *postRepo) DeleteCategoryIconDB(req *pbp.IdRequest) (*pbp.MessageResponse, error) {
	var res pbp.MessageResponse

	query := `
	DELETE FROM
		category_icons
	WHERE
		id = $1
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
		res.Message = "CategoryIcon not found for deletion."
	} else {
		res.Message = "CategoryIcon successfully deleted!"
	}

	return &res, nil
}

func (r *postRepo) ListCategoryIconsDB(req *pbp.GetAllRequest) (*pbp.ListCategoryIconResponse, error) {
	var allCategoryIcons pbp.ListCategoryIconResponse
	query := `
	SELECT
		id, 
		name,
		picture
	FROM 
		category_icons
	LIMIT $1
	OFFSET $2`
	offset := req.Limit * (req.Page - 1)
	rows, err := r.db.Query(query, req.Limit, offset)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var CategoryIcon pbp.CategoryIcon
		err := rows.Scan(
			&CategoryIcon.Id,
			&CategoryIcon.Name,
			&CategoryIcon.Picture)
		if err != nil {
			return nil, err
		}
		allCategoryIcons.Categoryicons = append(allCategoryIcons.Categoryicons, &CategoryIcon)
	}
	return &allCategoryIcons, nil
}
