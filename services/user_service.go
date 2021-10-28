package services

import (
	"github.com/yavuzsav/bookstore-users-api/domain/users"
	"github.com/yavuzsav/bookstore-users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	return nil, nil
}
