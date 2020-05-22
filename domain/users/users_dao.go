package users

import (
	"fmt"

	"github.com/rdelvallej32/bookstore_users-api/datasources/mysql/users_db"
	"github.com/rdelvallej32/bookstore_users-api/utils/errors"
	"github.com/rdelvallej32/bookstore_users-api/utils/mysql_utils"
)

const (
	queryInsertUser       = "INSERT INTO users(firstName, lastName, email, dateCreated, status, password) VALUES(?, ?, ?, ?, ?, ?);"
	queryGetUser          = "SELECT id, firstName, lastName, email, dateCreated FROM users WHERE id=?"
	queryUpdateUser       = "UPDATE users SET firstName=?, lastName=?, email=? WHERE id=?;"
	queryDeleteUser       = "DELETE FROM users WHERE id=?;"
	queryFindByUserStatus = "SELECT id, firstName, lastName, email, dateCreated, status FROM users WHERE status=?"
)

func (user *User) Get() *errors.RestErr {
	query, err := users_db.Client.Prepare(queryGetUser)

	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	defer query.Close()

	result := query.QueryRow(user.Id)

	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); getErr != nil {
		return mysql_utils.ParseError(getErr)
	}
	return nil
}

func (user *User) Save() *errors.RestErr {
	query, err := users_db.Client.Prepare(queryInsertUser)

	if err != nil {
		fmt.Println("ERROR With Statement")
		return errors.NewInternalServerError(err.Error())
	}

	defer query.Close()

	insertResult, saveErr := query.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, user.Password)

	if saveErr != nil {
		return mysql_utils.ParseError(saveErr)
	}

	userId, err := insertResult.LastInsertId()

	if err != nil {
		fmt.Println("ERROR with last insert id")
		return errors.NewInternalServerError(
			fmt.Sprintf("error when trying to save user: %s", err.Error()),
		)
	}

	user.Id = userId
	return nil
}

func (user *User) Update() *errors.RestErr {
	query, err := users_db.Client.Prepare(queryUpdateUser)

	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	defer query.Close()

	_, err = query.Exec(user.FirstName, user.LastName, user.Email, user.Id)

	if err != nil {
		return mysql_utils.ParseError(err)
	}

	return nil
}

func (user *User) Delete() *errors.RestErr {
	query, err := users_db.Client.Prepare(queryDeleteUser)

	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	defer query.Close()

	if _, err := query.Exec(user.Id); err != nil {
		return mysql_utils.ParseError(err)
	}

	return nil
}

func (user *User) FindByStatus(status string) ([]User, *errors.RestErr) {
	query, err := users_db.Client.Prepare(queryFindByUserStatus)

	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}

	defer query.Close()

	rows, err := query.Query(status)

	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}

	defer rows.Close()

	results := make([]User, 0)

	for rows.Next() {
		var user User

		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			return nil, mysql_utils.ParseError(err)
		}

		results = append(results, user)
	}

	if len(results) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("no users matching status %s", status))
	}

	return results, nil
}
