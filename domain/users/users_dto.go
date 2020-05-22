package users

import (
	"strings"

	"github.com/rdelvallej32/bookstore_users-api/utils/errors"
)

const (
	StatusActive = "active"
)

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	DateCreated string `json:"dateCreated"`
	Status      string `json:"status"`
	Password    string `json:"password"`
}

func (user *User) Validate() *errors.RestErr {
	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))

	if user.Email == "" {
		return errors.NewBadRequestError("invalid email address")
	}

	user.Password = strings.TrimSpace(user.Password)

	if user.Password == "" {
		return errors.NewBadRequestError("invalid password")
	}

	return nil
}
