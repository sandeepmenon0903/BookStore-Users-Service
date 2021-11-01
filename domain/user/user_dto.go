package user

import (
	"github.com/sandeepmenon0903/BookStore-Users-Service/utils/errors"
)

type User struct {
	UserId    int64  `json:"UserId"`
	UserName  string `json:"UserName"`
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Email     string `json:"Email"`
	CreatedOn string `json:"CreatedOn"`
	CreatedBy string `json:"CreatedBy"`
	Status    string `json:"status"`
	Password  string `json:"Password"` // This "-"  is given so that this doesnt include in the result
}

type Users []User

//Validate : This method is for validating the request user object.
func (user *User) Validate() *errors.RestErr {
	if user.Email == "" {
		if err := errors.BadRequest("Enter valid Email"); err != nil {
			return err
		}
	}
	return nil
}
