package user

import (
	"fmt"

	utils "github.com/sandeepmenon0903/BookStore-Users-Service/datasources/mysql/sql_utils"
	"github.com/sandeepmenon0903/BookStore-Users-Service/datasources/mysql/users_db"
	"github.com/sandeepmenon0903/BookStore-Users-Service/utils/errors"
)

const (
	querySearchByStatus = "select id, FirstName, LastName, Email, CreatedBy, CreatedOn,Status from users where status=?"
	queryDeleteUser     = "DELETE FROM USERS WHERE id=?"
	getNoRows           = "no rows in result set"
	queryGetUser        = "SELECT id, FirstName, LastName, Email, CreatedBy, CreatedOn,Status FROM users WHERE id=?;"
	queryInsertUser     = "INSERT INTO USERS(FirstName,LastName,Email,CreatedBy,CreatedOn,Status,Password)VALUES(?,?,?,?,?,?,?)"
	queryUpdateUser     = "Update USERS SET FirstName=?, LastName=?,Email=?,Status=? WHERE ID=?"
)

var (
	userDB = make(map[int64]*User)
)

// GetUser Function for getting the user based on userid.
func (user *User) GetUser() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		return utils.ParseError(err)
	}
	defer stmt.Close()
	result := stmt.QueryRow(user.UserId)
	if err := result.Scan(&user.UserId, &user.FirstName, &user.LastName, &user.Email, &user.CreatedBy, &user.CreatedOn, &user.Status); err != nil {
		return utils.ParseError(err)
	}
	return nil
}

// InsertUser function for inserting user.
func (user *User) InsertUser() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	result, err := stmt.Exec(user.UserName, user.LastName, user.Email, user.CreatedBy, user.CreatedOn, user.Status, user.Password)
	//Convert the error to a mysql error pointer
	if err != nil {
		return utils.ParseError(err)
	}
	userId, err := result.LastInsertId()
	if err != nil {
		return utils.ParseError(err)
	}
	user.UserId = userId
	return nil
}

//UpdateUser : Update the data for the user.
func (user *User) UpdateUser() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		return utils.ParseError(err)
	}
	defer stmt.Close()
	result, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.Status, user.UserId)
	if err != nil {
		return utils.ParseError(err)
	}
	fmt.Println(result)
	return nil
}

func (user *User) DeleteUser() (int64, *errors.RestErr) {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		return 0, utils.ParseError(err)
	}
	defer stmt.Close()
	result, err := stmt.Exec(user.UserId)
	if err != nil {
		return 0, utils.ParseError(err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, utils.ParseError(err)
	}
	return rowsAffected, nil
}

func (user *User) Search(status string) (Users, *errors.RestErr) {
	var users []User
	stmt, err := users_db.Client.Prepare(querySearchByStatus)
	if err != nil {
		return nil, utils.ParseError(err)
	}
	defer stmt.Close()
	result, err := stmt.Query(status)
	defer result.Close()
	if err != nil {
		return nil, utils.ParseError(err)
	}
	for result.Next() {
		var user_inst User
		if err := result.Scan(&user_inst.UserId, &user_inst.FirstName, &user_inst.LastName, &user_inst.Email, &user_inst.CreatedBy, &user_inst.CreatedOn, &user_inst.Status); err != nil {
			return nil, utils.ParseError(err)
		}
		users = append(users, user_inst)
		if len(users) == 0 {
			return nil, errors.NotFound("Cannot Find any Orders")
		}
	}
	return users, nil

}
