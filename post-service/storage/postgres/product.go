package postgres

import (
	"log"
	pbp "post-service/genproto/post"
)

// -------------------------------------------
// Product CRUD with Soft Delete
// -------------------------------------------

func (r *postRepo) CreateProductDB(req *pbp.Product) (*pbp.Product, error) {
	var res pbp.Product
	query := `
        INSERT INTO products(
            title,
            description,
            price,
			picture,
			category_id
        )
        VALUES ($1, $2, $3, $4, $5) 
        RETURNING 
            id,
            title,
            description,
            price,
			picture,
			category_id,
            created_at,
            updated_at
    `
	err := r.db.QueryRow(
		query,
		req.Title,
		req.Description,
		req.Price,
		req.Picture,
		req.CategoryId).Scan(
		&res.Id,
		&res.Title,
		&res.Description,
		&res.Price,
		&res.Picture,
		&res.CategoryId,
		&res.CreatedAt,
		&res.UpdatedAt)
	if err != nil {
		log.Println("err", err)
		return nil, err
	}
	return &res, nil
}

func (r *postRepo) ReadProductDB(req *pbp.IdRequest) (*pbp.Product, error) {
	var res pbp.Product
	query := `
        SELECT
			id,
			title,
			description,
			price,
			picture,
			category_id,
			created_at,
			updated_at
        FROM 
            products
        WHERE
            id=$1 AND
            deleted_at IS NULL
    `
	err := r.db.QueryRow(query, req.Id).Scan(
		&res.Id,
		&res.Title,
		&res.Description,
		&res.Price,
		&res.Picture,
		&res.CategoryId,
		&res.CreatedAt,
		&res.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (r *postRepo) UpdateProductDB(req *pbp.Product) (*pbp.Product, error) {
	var res pbp.Product
	query := `
        UPDATE
            products
        SET
            title=$1,
            description=$2,
            price=$3,
			picture=$4,
			category_id=$5,
            updated_at=CURRENT_TIMESTAMP
        WHERE
            id=$6
        RETURNING
            id,
            title,
            description,
            price,
            picture,
			category_id,
            created_at,
            updated_at
    `
	err := r.db.QueryRow(
		query,
		req.Title,
		req.Description,
		req.Price,
		req.Picture,
		req.CategoryId,
		req.Id).Scan(
		&res.Id,
		&res.Title,
		&res.Description,
		&res.Price,
		&res.Picture,
		&res.CategoryId,
		&res.CreatedAt,
		&res.UpdatedAt)
	if err != nil {
		log.Println("index:", req.Id)
		log.Println("err:", err)
		return nil, err
	}

	return &res, nil
}

func (r *postRepo) DeleteProductDB(req *pbp.IdRequest) (*pbp.MessageResponse, error) {
	var res pbp.MessageResponse

	query := `
        UPDATE
            products
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
		res.Message = "Product not found for deletion."
	} else {
		res.Message = "Product successfully deleted!"
	}

	return &res, nil
}

func (r *postRepo) ListProductsDB(req *pbp.GetAllRequest) (*pbp.ListProductResponse, error) {
	var allProducts pbp.ListProductResponse
	query := `
        SELECT
			id,
			title,
			description,
			price,
			picture,
			category_id,
			created_at,
			updated_at
        FROM 
            products
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
		var Product pbp.Product
		err := rows.Scan(
			&Product.Id,
			&Product.Title,
			&Product.Description,
			&Product.Price,
			&Product.Picture,
			&Product.CategoryId,
			&Product.CreatedAt,
			&Product.UpdatedAt)
		if err != nil {
			return nil, err
		}
		allProducts.Products = append(allProducts.Products, &Product)
	}
	return &allProducts, nil
}
