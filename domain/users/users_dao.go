package users

import (
	"fmt"

	"github.com/rdelvallej32/bookstore_users-api/datasources/mysql/users_db"
	"github.com/rdelvallej32/bookstore_users-api/logger"
	"github.com/rdelvallej32/bookstore_users-api/utils/errors"
)

const (
	queryInsertUser       = "INSERT INTO users(firstName, lastName, email, dateCreated, status, password) VALUES(?, ?, ?, ?, ?, ?);"
	queryGetUser          = "SELECT id, firstName, lastName, email, dateCreated, status FROM users WHERE id=?"
	queryUpdateUser       = "UPDATE users SET firstName=?, lastName=?, email=? WHERE id=?;"
	queryDeleteUser       = "DELETE FROM users WHERE id=?;"
	queryFindByUserStatus = "SELECT id, firstName, lastName, email, dateCreated, status FROM users WHERE status=?"
)

func (user *User) Get() *errors.RestErr {
	query, err := users_db.Client.Prepare(queryGetUser)

	if err != nil {
		logger.Error("error preparing get user statement", err)
		return errors.NewInternalServerError("database error")
	}

	defer query.Close()

	result := query.QueryRow(user.Id)

	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); getErr != nil {
		logger.Error("error when trying to get user by id", err)
		return errors.NewInternalServerError("database error")
	}
	return nil
}

func (user *User) Save() *errors.RestErr {
	query, err := users_db.Client.Prepare(queryInsertUser)

	if err != nil {
		logger.Error("error preparing insert user statement", err)
		return errors.NewInternalServerError("database error")
	}

	defer query.Close()

	insertResult, saveErr := query.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, user.Password)

	if saveErr != nil {
		logger.Error("error when trying to save user", err)
		return errors.NewInternalServerError("database error")
	}

	userId, err := insertResult.LastInsertId()

	if err != nil {
		logger.Error("error when trying to get last user id  after creating a new user", err)
		return errors.NewInternalServerError("database error")
	}

	user.Id = userId
	return nil
}

func (user *User) Update() *errors.RestErr {
	query, err := users_db.Client.Prepare(queryUpdateUser)

	if err != nil {
		logger.Error("error preparing update user statement", err)
		return errors.NewInternalServerError("database error")
	}

	defer query.Close()

	_, err = query.Exec(user.FirstName, user.LastName, user.Email, user.Id)

	if err != nil {
		logger.Error("error updating user", err)
		return errors.NewInternalServerError("database error")
	}

	return nil
}

func (user *User) Delete() *errors.RestErr {
	query, err := users_db.Client.Prepare(queryDeleteUser)

	if err != nil {
		logger.Error("error preparing delete user statement", err)
		return errors.NewInternalServerError("database error")
	}

	defer query.Close()

	if _, err := query.Exec(user.Id); err != nil {
		logger.Error("error deleting user", err)
		return errors.NewInternalServerError("database error")
	}

	return nil
}

func (user *User) FindByStatus(status string) ([]User, *errors.RestErr) {
	query, err := users_db.Client.Prepare(queryFindByUserStatus)

	if err != nil {
		logger.Error("error preparing user status statement", err)
		return nil, errors.NewInternalServerError("database error")
	}

	defer query.Close()

	rows, err := query.Query(status)

	if err != nil {
		logger.Error("error finding user by status", err)
		return nil, errors.NewInternalServerError("database error")
	}

	defer rows.Close()

	results := make([]User, 0)

	for rows.Next() {
		var user User

		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			logger.Error("error scanning user row into user struct", err)
			return nil, errors.NewInternalServerError("database error")
		}

		results = append(results, user)
	}

	if len(results) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("no users matching status %s", status))
	}

	return results, nil
}
