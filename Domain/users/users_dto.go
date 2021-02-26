package users

import (
	"strings"

	"github.com/projects/REST-GO-GIN/Utils/erros"
)

// User .......
type User struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

// Validate ..
func (user *User) Validate() *erros.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return erros.BadRequestError("Invalid email address")
	}
	user.Password = strings.TrimSpace(user.Password)
	if user.Password == "" {
		return erros.BadRequestError("Invalid password")
	}
	return nil
}
