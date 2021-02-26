package services

import (
	"github.com/projects/REST-GO-GIN/Domain/users"
	"github.com/projects/REST-GO-GIN/Utils/erros"
)

// CreateUser ..
func CreateUser(user users.User) (*users.User, *erros.RestErr) {
	// Validate
	if err := user.Validate(); err != nil {
		return nil, err
	}

	//Save to database
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil

}

//GetUser ..
func GetUser(userID int64) (*users.User, *erros.RestErr) {

	result := &users.User{ID: userID}

	if err := result.Get(); err != nil {
		return nil, err
	}

	return result, nil

}
