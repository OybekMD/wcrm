package postgres

import (
	"errors"
	"strconv"
	pbu "user-service/genproto/user"
)

/*
QueryBuildCreateUser constructs an SQL INSERT query and parameter list based on the provided User protobuf message for creating a new user.

Parameters:

	reqData pbu.User: A protobuf User message containing the data for creating a new user.

Returns:

	string: The generated SQL INSERT query.
	[]interface{}: A slice of interface{} containing the parameters for the query.

Description:
This function takes a User protobuf message and constructs an SQL INSERT query to insert a new user record into the database. It dynamically builds the query based on the fields in the User message, adding columns and values as needed. The function returns the generated SQL query string and a slice of interface{} containing the parameter values.

Note:
- The function assumes the availability of the `pbu` package which contains the definition of the User protobuf message.
- It uses positional parameters in the SQL query to avoid SQL injection vulnerabilities.
- The returned query includes placeholders for columns and values and does not execute the query directly.
*/
func QueryBuildCreateUser(reqData pbu.User) (string, []interface{}) {
	args := []interface{}{
		reqData.Id,
		reqData.FirstName,
		reqData.LastName,
	}
	query := "INSERT INTO users(id, first_name, last_name"
	reply := " RETURNING id, first_name, last_name"
	count := 3
	if reqData.Username != "" {
		query += ", username"
		reply += ", username"
		args = append(args, reqData.Username)
		count++
	}
	if reqData.PhoneNumber != "" {
		query += ", phone_number"
		reply += ", phone_number"
		args = append(args, reqData.PhoneNumber)
		count++
	}
	if reqData.Bio != "" {
		query += ", bio"
		reply += ", bio"
		args = append(args, reqData.Bio)
		count++
	}
	if reqData.BirthDay != "" {
		query += ", birth_day"
		reply += ", birth_day"
		args = append(args, reqData.BirthDay)
		count++
	}

	query += ", email"
	reply += ", email"
	args = append(args, reqData.Email)
	count++

	if reqData.Avatar != "" {
		query += ", avatar"
		reply += ", avatar"
		args = append(args, reqData.Avatar)
		count++
	}

	query += ", password"
	reply += ", password"
	args = append(args, reqData.Password)
	count++

	query += ", refresh_token"
	args = append(args, reqData.RefreshToken)
	count++

	reply += ", refresh_token"

	query += ") VALUES ("
	for i := 1; i <= count; i++ {
		num := strconv.Itoa(i)
		if i == count {
			query = query + "$" + num + ")"
		} else {
			query = query + "$" + num + ", "
		}
	}
	reply += ", created_at, updated_at"
	query += reply

	return query, args
}

/*
QueryBuildPatchUser constructs an SQL UPDATE query and parameter list based on the provided User protobuf message.

Parameters:

	req *pbu.User: A protobuf User message containing the data to update.

Returns:

	string: The generated SQL UPDATE query.
	[]interface{}: A slice of interface{} containing the parameters for the query.
	error: An error, if any occurred during the process.

Description:
This function takes a User protobuf message and constructs an SQL UPDATE query to update the corresponding user record in the database. It dynamically builds the query based on the non-empty fields in the User message, utilizing positional parameters for the values. The function returns the generated SQL query string, a slice of interface{} containing the parameter values, and an error if there are no fields to update.

Note:
- The function assumes the availability of the `pbu` package which contains the definition of the User protobuf message.
- It uses positional parameters in the SQL query to avoid SQL injection vulnerabilities.
*/
func QueryBuildPatchUser(req *pbu.User) (string, []interface{}, error) {
	var query string
	var args []interface{}
	query = "UPDATE users SET "

	updatedFields := false

	step := 1
	if req.FirstName != "" {
		query += "first_name = $" + strconv.Itoa(step) + ", "
		step++
		args = append(args, req.FirstName)
		updatedFields = true
	}
	if req.LastName != "" {
		query += "last_name = $" + strconv.Itoa(step) + ", "
		step++
		args = append(args, req.LastName)
		updatedFields = true
	}
	if req.Username != "" {
		query += "username = $" + strconv.Itoa(step) + ", "
		step++
		args = append(args, req.Username)
		updatedFields = true
	}

	if req.PhoneNumber != "" {
		query += "phone_number = $" + strconv.Itoa(step) + ", "
		step++
		args = append(args, req.PhoneNumber)
		updatedFields = true
	}

	if req.Bio != "" {
		query += "bio = $" + strconv.Itoa(step) + ", "
		step++
		args = append(args, req.Bio)
		updatedFields = true
	}

	if req.BirthDay != "" {
		query += "birth_day = $" + strconv.Itoa(step) + ", "
		step++
		args = append(args, req.BirthDay)
		updatedFields = true
	}

	if req.Email != "" {
		query += "email = $" + strconv.Itoa(step) + ", "
		step++
		args = append(args, req.Email)
		updatedFields = true
	}

	if req.Avatar != "" {
		query += "avatar = $" + strconv.Itoa(step) + ", "
		step++
		args = append(args, req.Avatar)
		updatedFields = true
	}

	if req.Password != "" {
		query += "password = $" + strconv.Itoa(step) + ", "
		step++
		args = append(args, req.Password)
		updatedFields = true
	}

	if !updatedFields {
		return "", nil, errors.New("no fields to update")
	}
	query += "updated_at=CURRENT_TIMESTAMP WHERE id=$" + strconv.Itoa(step)
	args = append(args, req.Id)

	return query, args, nil
}
