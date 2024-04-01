package postgres

import (
    pbp "post-service/genproto/post"
)

// -------------------------------------------
// OrderProduct CRUD
// -------------------------------------------

func (r *postRepo) CreateOrderproductDB(req *pbp.Orderproduct) (*pbp.Orderproduct, error) {
    var res pbp.Orderproduct
    query := `
        INSERT INTO orderproducts(
            user_id,
            product_id
        )
        VALUES ($1, $2) 
        RETURNING 
            id,
            user_id,
            product_id,
            created_at
    `
    err := r.db.QueryRow(
        query,
        req.UserId,
        req.ProductId).Scan(
        &res.Id,
        &res.UserId,
        &res.ProductId,
        &res.CreatedAt)
    if err != nil {
        return nil, err
    }
    return &res, nil
}

func (r *postRepo) ReadOrderproductDB(req *pbp.IdRequest) (*pbp.Orderproduct, error) {
    var res pbp.Orderproduct
    query := `
        SELECT
            id,
            user_id,
            product_id,
            created_at
        FROM 
            orderproducts
        WHERE
            id=$1
    `
    err := r.db.QueryRow(query, req.Id).Scan(
        &res.Id,
        &res.UserId,
        &res.ProductId,
        &res.CreatedAt)
    if err != nil {
        return nil, err
    }

    return &res, nil
}

func (r *postRepo) UpdateOrderproductDB(req *pbp.Orderproduct) (*pbp.Orderproduct, error) {
    var res pbp.Orderproduct
    query := `
        UPDATE
            orderproducts
        SET
            user_id=$1,
            product_id=$2
        WHERE
            id=$3
        RETURNING
            id,
            user_id,
            product_id,
            created_at
    `
    err := r.db.QueryRow(
        query,
        req.UserId,
        req.ProductId,
        req.Id).Scan(
        &res.Id,
        &res.UserId,
        &res.ProductId,
        &res.CreatedAt)
    if err != nil {
        return nil, err
    }

    return &res, nil
}

func (r *postRepo) DeleteOrderproductDB(req *pbp.IdRequest) (*pbp.MessageResponse, error) {
    var res pbp.MessageResponse

    query := `
        DELETE FROM
            orderproducts
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
        res.Message = "Orderproduct not found for deletion."
    } else {
        res.Message = "Orderproduct successfully deleted!"
    }

    return &res, nil
}

func (r *postRepo) ListOrderproductsDB(req *pbp.GetAllRequest) (*pbp.ListOrderproductResponse, error) {
    var allOrderProducts pbp.ListOrderproductResponse
    query := `
        SELECT
            id, 
            user_id,
            product_id,
            created_at
        FROM 
            orderproducts
        LIMIT $1
        OFFSET $2
    `
    offset := req.Limit * (req.Page - 1)
    rows, err := r.db.Query(query, req.Limit, offset)
    if err != nil {
        return nil, err
    }
    for rows.Next() {
        var Orderproduct pbp.Orderproduct
        err := rows.Scan(
            &Orderproduct.Id,
            &Orderproduct.UserId,
            &Orderproduct.ProductId,
            &Orderproduct.CreatedAt)
        if err != nil {
            return nil, err
        }
        allOrderProducts.Orderproducts = append(allOrderProducts.Orderproducts, &Orderproduct)
    }
    return &allOrderProducts, nil
}
