package users

import (
	"fmt"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/rdelvallej32/bookstore_users-api/datasources/mysql/users_db"
	"github.com/rdelvallej32/bookstore_users-api/utils/date_util"
	"github.com/rdelvallej32/bookstore_users-api/utils/errors"
)

const (
	errorNoRows     = "no rows in result set"
	queryInsertUser = "INSERT INTO users(firstName, lastName, email, dateCreated) VALUES(?, ?, ?, ?);"
	queryGetUser    = "SELECT id, firstName, lastName, email, dateCreated FROM users WHERE id=?"
)

func (user *User) Get() *errors.RestErr {
	query, err := users_db.Client.Prepare(queryGetUser)

	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	defer query.Close()

	result := query.QueryRow(user.Id)

	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
		}

		return errors.NewInternalServerError(
			fmt.Sprintf("error when trying to get user %d: %s", user.Id, err.Error()),
		)
	}

	// user.Id = result.Id
	// user.FirstName = result.FirstName
	// user.LastName = result.LastName
	// user.Email = result.Email
	// user.DateCreated = result.DateCreated
	return nil
}

func (user *User) Save() *errors.RestErr {
	query, err := users_db.Client.Prepare(queryInsertUser)

	if err != nil {
		fmt.Println("ERROR With Statement")
		return errors.NewInternalServerError(err.Error())
	}

	defer query.Close()

	user.DateCreated = date_util.GetNowString()

	insertResult, saveErr := query.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)

	if saveErr != nil {
		sqlErr, ok := saveErr.(*mysql.MySQLError)

		if !ok {
			return errors.NewInternalServerError(
				fmt.Sprintf("error when trying to save user: %s", saveErr.Error()),
			)
		}
		switch sqlErr.Number {
		case 1062:
			return errors.NewBadRequestError(fmt.Sprintf("email %s already exists", user.Email))
		}

		return errors.NewInternalServerError(fmt.Sprintf("error when trying to save user: %s", saveErr.Error()))
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
