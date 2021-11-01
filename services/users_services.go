package services

import (
	"fmt"
	"strings"

	date_utils "github.com/sandeepmenon0903/BookStore-Users-Service/date-utils"
	"github.com/sandeepmenon0903/BookStore-Users-Service/domain/user"
	"github.com/sandeepmenon0903/BookStore-Users-Service/utils/crypto"
	"github.com/sandeepmenon0903/BookStore-Users-Service/utils/errors"
)

const (
	statusActive = "Active"
)

var (
	UsersService userServiceInterface = &usersService{}
)

type usersService struct{}

type userServiceInterface interface {
	CreateUser(user user.User) (*user.User, *errors.RestErr)
	GetUser(userId int64) (*user.User, *errors.RestErr)
	UpdateUser(user user.User, partial bool) *errors.RestErr
	DeleteUser(userId int64) *errors.RestErr
	SearchUsers(status string) (user.Users, *errors.RestErr)
}

// CreateUser : This service will be invoking the call to the database.
// The user object or the error will be returned.
func (service *usersService) CreateUser(user user.User) (*user.User, *errors.RestErr) {
	if strings.TrimSpace(user.Status) == "Active" {
		user.Status = statusActive
	}
	user.CreatedOn = date_utils.GetNowDbFormat()
	if user.Password != "" {
		user.Password = crypto.GetMd5(user.Password)
	}
	fmt.Println(date_utils.TimeNow())
	if err := user.InsertUser(); err != nil {
		return nil, err
	}
	return &user, nil
}

func (service *usersService) GetUser(userId int64) (*user.User, *errors.RestErr) {
	if userId <= 0 {
		return nil, errors.BadRequest("User Id should be valid value")
	}
	var userDetails user.User
	userDetails.UserId = userId
	if err := userDetails.GetUser(); err != nil {
		return nil, err
	}
	return &userDetails, nil
}

func (service *usersService) UpdateUser(user user.User, partial bool) *errors.RestErr {
	currentUser := user
	currentUser.UserId = user.UserId
	if err := currentUser.GetUser(); err != nil {
		return err
	}
	// This change is to support the patch request
	if partial {
		if strings.TrimSpace(user.FirstName) == "" {
			user.FirstName = currentUser.FirstName
		}
		if strings.TrimSpace(user.LastName) == "" {
			user.LastName = currentUser.LastName
		}
		if strings.TrimSpace(user.UserName) == "" {
			user.UserName = currentUser.UserName
		}
		if strings.TrimSpace(user.Status) == "" {
			user.Status = currentUser.Status
		}
	}

	err := user.UpdateUser()
	if err != nil {
		return err
	}
	return nil
}

func (service *usersService) DeleteUser(userId int64) *errors.RestErr {
	var userDetails user.User
	userDetails.UserId = userId
	if err := userDetails.GetUser(); err != nil {
		return err
	}
	rowsAffected, err := userDetails.DeleteUser()
	if err != nil {
		return err
	}
	if rowsAffected <= 0 {
		return errors.NewInternalServerError(fmt.Sprintf("Could Delete the User %s", userDetails.FirstName))
	}
	return nil
}

// Service for Searching the user based on the status
func (service *usersService) SearchUsers(status string) (user.Users, *errors.RestErr) {
	var userDetails user.User
	users, err := userDetails.Search(status)
	if err != nil {
		return nil, err
	}
	return users, err
}
