package users

import (

	//"github.com/lib/pq" ..
	_ "github.com/lib/pq"

	"github.com/projects/REST-GO-GIN/Datasource/usersdb"
	"github.com/projects/REST-GO-GIN/Utils/erros"
)

const (
	queryInserUser = "INSERT INTO accounts(first_name, last_name, email, password) VALUES (?, ?, ?, ?);"
	queryGetUser   = "SELECT user_id, first_name, last_name, email FROM accounts WHERE user_id=?;"
)

// Get ..
func (user *User) Get() *erros.RestErr {
	stmt, err := usersdb.Client.Prepare(queryGetUser)
	if err != nil {
		return erros.InternalServerError("Database error")
	}

	defer stmt.Close()

	result := stmt.QueryRow(user.ID)
	if getErr := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email); getErr != nil {
		return erros.InternalServerError("database error ")
	}

	return nil

}

// Save ..
func (user *User) Save() *erros.RestErr {

	stmt, err := usersdb.Client.Prepare(queryInserUser)

	if err != nil {
		return erros.InternalServerError("database error")
	}

	defer stmt.Close()

	inserResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.Password)
	if saveErr != nil {
		return erros.InternalServerError("database error")
	}

	userID, err := inserResult.LastInsertId()
	if err != nil {
		return erros.InternalServerError("database error")
	}
	user.ID = userID
	return nil
}
