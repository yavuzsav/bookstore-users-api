package users

import "github.com/yavuzsav/bookstore-users-api/utils/errors"

func (user User) Get() (*User, *errors.RestErr) {
	return nil, nil
}

func (user User) Save() *errors.RestErr {
	return nil
}