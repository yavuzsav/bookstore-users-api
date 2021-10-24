package services

import (
	"github.com/yavuzsav/bookstore-users-api/domain/users"
	"github.com/yavuzsav/bookstore-users-api/utils/errors"
	"net/http"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	//return &user, nil


	return &user, &errors.RestErr{
		Status:  http.StatusInternalServerError,
	}
}
