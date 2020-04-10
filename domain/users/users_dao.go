package users

import (
	"fmt"
	"time"

	"github.com/rdelvallej32/bookstore_users-api/datasources/mysql/users_db"
	"github.com/rdelvallej32/bookstore_users-api/utils/errors"
)

var (
	usersDb = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {
	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}

	result := usersDb[user.Id]

	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
	}

	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated
	return nil
}

func (user *User) Save() *errors.RestErr {
	current := usersDb[user.Id]
	if current != nil {

		if current.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already registerd", user.Email))
		}

		return errors.NewBadRequestError(fmt.Sprintf("user %d already exists", user.Id))
	}

	now := time.Now().UTC()
	user.DateCreated = now.Format("2006-01-02T15:04:05Z")

	usersDb[user.Id] = user
	return nil
}
