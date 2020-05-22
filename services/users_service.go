package services

import (
	"github.com/rdelvallej32/bookstore_users-api/domain/users"
	"github.com/rdelvallej32/bookstore_users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUser(userId int64) (*users.User, *errors.RestErr) {
	result := &users.User{
		Id: userId,
	}

	if err := result.Get(); err != nil {
		return nil, err
	}

	return result, nil
}

func UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr) {
	currentUser, err := GetUser(user.Id)

	if err != nil {
		return nil, err
	}

	if isPartial {
		if user.FirstName != "" {
			currentUser.FirstName = user.FirstName
		}

		if user.LastName != "" {
			currentUser.LastName = user.LastName
		}

		if user.Email != "" {
			currentUser.Email = user.Email
		}
	} else {
		currentUser.FirstName = user.FirstName
		currentUser.LastName = user.LastName
		currentUser.Email = user.Email
	}

	if err := currentUser.Update(); err != nil {
		return nil, err
	}

	return currentUser, nil
}
