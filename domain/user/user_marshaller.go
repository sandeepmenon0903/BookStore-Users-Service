package user

import (
	"encoding/json"
	"fmt"
)

type PublicUser struct {
	Email     string `json:"Email"`
	CreatedOn string `json:"CreatedOn"`
	CreatedBy string `json:"CreatedBy"`
	Status    string `json:"status"`
}

type PrivateUser struct {
	UserId    int64  `json:"UserId"`
	UserName  string `json:"UserName"`
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Email     string `json:"Email"`
	CreatedOn string `json:"CreatedOn"`
	CreatedBy string `json:"CreatedBy"`
	Status    string `json:"status"`
}

func (user *User) Marshall(isPublic bool) interface{} {
	if isPublic {
		return PublicUser{
			Email:     user.Email,
			CreatedOn: user.CreatedOn,
			CreatedBy: user.CreatedBy,
			Status:    user.Status,
		}
	}
	userJson, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Errored while processing the request")
	}
	var internalUser PrivateUser
	json.Unmarshal(userJson, &internalUser)
	return internalUser
}

func (users Users) Marshall(isPublic bool) []interface{} {
	result := make([]interface{}, len(users))
	for index, user := range users {
		result[index] = user.Marshall(isPublic)
	}
	return result
}
