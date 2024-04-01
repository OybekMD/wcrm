package postgres

import (
	"database/sql"
	pbu "user-service/genproto/user"

	"github.com/jmoiron/sqlx"
)

type userRepo struct {
	db *sqlx.DB
}

// NewUserRepo ...c
func NewUserRepo(db *sqlx.DB) *userRepo {
	return &userRepo{db: db}
}

// -------------------------------------------
// USER CRUD
// -------------------------------------------

func (r *userRepo) CreateUserDB(req *pbu.User) (*pbu.User, error) {
	var res pbu.User
	var nullUsername, nullPhoneNumber, nullBio, nullBirthDay, nullAvatar sql.NullString

	query, args := QueryBuildCreateUser(*req)

	err := r.db.QueryRow(
		query, args...).Scan(
		&res.Id,
		&res.FirstName,
		&res.LastName,
		&nullUsername,
		&nullPhoneNumber,
		&nullBio,
		&nullBirthDay,
		&res.Email,
		&nullAvatar,
		&res.Password,
		&res.RefreshToken,
		&res.CreatedAt,
		&res.UpdatedAt)
	if err != nil {
		return nil, err
	}

	if nullUsername.Valid {
		res.Username = nullUsername.String
	}

	if nullPhoneNumber.Valid {
		res.PhoneNumber = nullPhoneNumber.String
	}

	if nullBio.Valid {
		res.Bio = nullBio.String
	}

	if nullBirthDay.Valid {
		res.BirthDay = nullBirthDay.String
	}

	if nullAvatar.Valid {
		res.Avatar = nullAvatar.String
	}

	return &res, nil
}

func (r *userRepo) ReadUserDB(req *pbu.IdRequest) (*pbu.User, error) {
	var user pbu.User
	var nullUsername, nullPhoneNumber, nullBio, nullBirthDay, nullAvatar sql.NullString
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
		id=$1 
	AND deleted_at IS NULL
	`
	err := r.db.QueryRow(query, req.Id).Scan(
		&user.Id,
		&user.FirstName,
		&user.LastName,
		&nullUsername,
		&nullPhoneNumber,
		&nullBio,
		&nullBirthDay,
		&user.Email,
		&nullAvatar,
		&user.Password,
		&user.RefreshToken,
		&user.CreatedAt,
		&user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	if nullUsername.Valid {
		user.Username = nullUsername.String
	}

	if nullPhoneNumber.Valid {
		user.PhoneNumber = nullPhoneNumber.String
	}

	if nullBio.Valid {
		user.Bio = nullBio.String
	}

	if nullBirthDay.Valid {
		user.BirthDay = nullBirthDay.String
	}

	if nullAvatar.Valid {
		user.Avatar = nullAvatar.String
	}

	return &user, nil
}

// PUT
func (r *userRepo) UpdateUserDB(req *pbu.User) (*pbu.User, error) {
	var res pbu.User
	query := `
	UPDATE
		users
	SET
		first_name=$1, 
		last_name=$2, 
		username=$3, 
		phone_number=$4,
		bio=$5,
		birth_day=$6,
		email=$7,
		avatar=$8,
		updated_at=CURRENT_TIMESTAMP
	WHERE
		id=$9
	RETURNING
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
	`
	err := r.db.QueryRow(
		query,
		req.FirstName,
		req.LastName,
		req.Username,
		req.PhoneNumber,
		req.Bio,
		req.BirthDay,
		req.Email,
		req.Avatar,
		req.Id).Scan(
		&res.Id,
		&res.FirstName,
		&res.LastName,
		&res.Username,
		&res.PhoneNumber,
		&res.Bio,
		&res.BirthDay,
		&res.Email,
		&res.Avatar,
		&res.Password,
		&res.RefreshToken,
		&res.CreatedAt,
		&res.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// PATCH
func (r *userRepo) PatchUpdateUserDB(req *pbu.User) (*pbu.User, error) {
	var res pbu.User
	var nullUsername, nullPhoneNumber, nullBio, nullBirthDay, nullAvatar sql.NullString
	query, arg, err := QueryBuildPatchUser(req)
	if err != nil {
		return nil, err
	}
	query += `
	RETURNING
		id, 
		first_name, 
		last_name,
		username, 
		bio,
		birth_day,
		email,
		password_hash,
		avatar,
		experience_level,
		coint,
		score,
		refresh_token,
		created_at,
		updated_at
	`
	err = r.db.QueryRow(
		query,
		arg...).Scan(
		&res.Id,
		&res.FirstName,
		&res.LastName,
		&nullUsername,
		&nullPhoneNumber,
		&nullBio,
		&nullBirthDay,
		&res.Email,
		&nullAvatar,
		&res.RefreshToken,
		&res.CreatedAt,
		&res.UpdatedAt)
	if err != nil {
		return nil, err
	}
	if nullUsername.Valid {
		res.Username = nullUsername.String
	}
	if nullBio.Valid {
		res.Bio = nullBio.String
	}
	if nullBirthDay.Valid {
		res.BirthDay = nullBirthDay.String
	}
	if nullAvatar.Valid {
		res.Avatar = nullAvatar.String
	}

	return &res, nil
}

func (r *userRepo) DeleteUserDB(req *pbu.IdRequest) (*pbu.MessageResponse, error) {
	var res pbu.MessageResponse

	query := `
	UPDATE
		users
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
		res.Message = "User not found for deletion."
	} else {
		res.Message = "User successfully deleted!"
	}

	return &res, nil
}

func (r *userRepo) ListUserDB(req *pbu.GetAllRequest) (*pbu.ListUserResponse, error) {
	var allUser pbu.ListUserResponse
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
		deleted_at IS NULL
	LIMIT $1
	OFFSET $2`
	offset := req.Limit * (req.Page - 1)
	rows, err := r.db.Query(query, req.Limit, offset)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var user pbu.User
		var nullUsername, nullPhoneNumber, nullBio, nullBirthDay, nullAvatar sql.NullString
		err := rows.Scan(
			&user.Id,
			&user.FirstName,
			&user.LastName,
			&nullUsername,
			&nullPhoneNumber,
			&nullBio,
			&nullBirthDay,
			&user.Email,
			&nullAvatar,
			&user.Password,
			&user.RefreshToken,
			&user.CreatedAt,
			&user.UpdatedAt)
		if err != nil {
			return nil, err
		}

		if nullUsername.Valid {
			user.Username = nullUsername.String
		}

		if nullPhoneNumber.Valid {
			user.PhoneNumber = nullPhoneNumber.String
		}

		if nullBio.Valid {
			user.Bio = nullBio.String
		}

		if nullBirthDay.Valid {
			user.BirthDay = nullBirthDay.String
		}

		if nullAvatar.Valid {
			user.Avatar = nullAvatar.String
		}

		allUser.Users = append(allUser.Users, &user)
	}
	return &allUser, nil
}
