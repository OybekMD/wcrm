package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	pbu "user-service/genproto/user"
	"user-service/pkg/etc"
)

func (r *userRepo) LoginDB(req *pbu.LoginRequest) (*pbu.User, error) {
	var user pbu.User
	var nulluname, nullphonenumber, nullbday, nullAva, nullBio sql.NullString

	query := `
	SELECT
		id, 
		first_name, 
		last_name,
		username, 
		phone_number,
		bio,
		birth_day,
		email,
		avatar,
		password,
		refresh_token,
		created_at,
		updated_at
	FROM 
		users
	WHERE
		email=$1 OR
		username = $1
	AND deleted_at IS NULL
	`
	err := r.db.QueryRow(query, req.Email).Scan(
		&user.Id,
		&user.FirstName,
		&user.LastName,
		&nulluname,
		&nullphonenumber,
		&nullBio,
		&nullbday,
		&user.Email,
		&nullAva,
		&user.Password,
		&user.RefreshToken,
		&user.CreatedAt,
		&user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	if nulluname.Valid {
		user.Username = nulluname.String
	}
	if nullphonenumber.Valid {
		user.PhoneNumber = nullphonenumber.String
	}
	if nullBio.Valid {
		user.Bio = nullBio.String
	}
	if nullbday.Valid {
		user.BirthDay = nullbday.String
	}
	if nullAva.Valid {
		user.Avatar = nullAva.String
	}
	if etc.CheckPasswordHash(req.Password, user.Password) {
		return &user, nil
	}
	return nil, errors.New("password is not correct")
}

func (r *userRepo) CheckUniqueDB(req *pbu.CheckUniqueRequest) (*pbu.CheckUniqueRespons, error) {
	query := fmt.Sprintf("SELECT count(1) FROM users WHERE %s = $1 AND deleted_at IS NULL", req.Column)
	var result int

	err := r.db.QueryRow(query, req.Value).Scan(&result)
	if err != nil {
		log.Println("Error Scan CheckField")
		return nil, err
	}

	if result == 0 {
		return &pbu.CheckUniqueRespons{
			IsExist: false,
		}, nil
	}

	return &pbu.CheckUniqueRespons{IsExist: true}, nil
}

func (r *userRepo) UpdatePasswordDB(req *pbu.UpdatePasswordRequest) (*pbu.MessageResponse, error) {
	var res pbu.MessageResponse
	query := `
		UPDATE
			users
		SET
			password=$1
		WHERE
			email=$2
		`

	_, err := r.db.Query(query, req.Password, req.Email)
	if err != nil {
		fmt.Println("error while updating user password!")
		// log.Println("error while updating user password!", err.Error())
		return nil, err
	} else {
		res.Message = "Password Succesfully updated!"
	}

	return &res, nil
}


func (r *userRepo) GetFullNameDB(req *pbu.LoginRequest) (*pbu.User, error) {
	var user pbu.User
	query := `
	SELECT
		first_name, 
		last_name
	FROM 
		users
	WHERE
		email=$1
	AND deleted_at IS NULL
	`
	err := r.db.QueryRow(query, req.Email).Scan(
		&user.FirstName,
		&user.LastName)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepo) IsAdminDB(req *pbu.IdRequest) (*pbu.CheckUniqueRespons, error) {
	var res pbu.CheckUniqueRespons
	query := `
	SELECT
		id
	FROM 
		admins
	WHERE
		email=$1
	AND deleted_at IS NULL
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
		res.IsExist = false
	} else {
		res.IsExist = true
	}

	return &res, nil
}