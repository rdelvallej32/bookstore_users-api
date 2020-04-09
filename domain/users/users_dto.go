package users

import (
	"strings"

	"github.com/rdelvallej32/bookstore_users-api/utils/date_util"
	"github.com/rdelvallej32/bookstore_users-api/utils/errors"
)

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	DateCreated string `json:"dateCreated"`
}

func (user *User) Validate() *errors.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))

	if user.Email == "" {
		return errors.NewBadRequestError("invalid email address")
	}

	user.DateCreated = date_util.GetNowString()
	return nil
}
