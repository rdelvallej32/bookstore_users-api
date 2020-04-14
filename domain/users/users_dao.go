package users

import (
	"fmt"

	"github.com/rdelvallej32/bookstore_users-api/datasources/mysql/users_db"
	"github.com/rdelvallej32/bookstore_users-api/utils/date_util"
	"github.com/rdelvallej32/bookstore_users-api/utils/errors"
	"github.com/rdelvallej32/bookstore_users-api/utils/mysql_utils"
)

const (
	queryInsertUser = "INSERT INTO users(firstName, lastName, email, dateCreated) VALUES(?, ?, ?, ?);"
	queryGetUser    = "SELECT id, firstName, lastName, email, dateCreated FROM users WHERE id=?"
	queryUpdateUser = "UPDATE users SET firstName=?, lastName=?, email=? WHERE id=?;"
)

func (user *User) Get() *errors.RestErr {
	query, err := users_db.Client.Prepare(queryGetUser)

	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	defer query.Close()

	result := query.QueryRow(user.Id)

	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); getErr != nil {
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

	user.DateCreated = date_util.GetNowString()

	insertResult, saveErr := query.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)

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
