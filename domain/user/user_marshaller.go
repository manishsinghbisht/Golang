package user

import (
	// "database/sql"
	"encoding/json"
	// "gopkg.in/guregu/null.v4"
)

type PublicUser struct {
	Id          int64  `json:"Id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Gender      string `json:"gender"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	PostalPin      int64 `json:"postal_pin"`
}

type PrivateUser struct {
	Id          int64  `json:"Id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Gender      string `json:"gender"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	Password    string `json:"password"`
	Mobile      int64 `json:"mobile"`
	MobilePin      int64 `json:"mobile_pin"`
	PostalPin      int64 `json:"postal_pin"`
}

func (users Users) Marshall(isPublic bool) []interface{} {
	result := make([]interface{}, len(users))
	for index, user := range users {
		result[index] = user.Marshall(isPublic)
	}
	return result
}

func (user *User) Marshall(isPublic bool) interface{} {
	if isPublic {
		return PublicUser{
			Id:          user.Id,
			DateCreated: user.DateCreated,
			Status:      user.Status,
			PostalPin: user.PostalPin,
		}
	}

	userJson, _ := json.Marshal(user)
	var privateUser PrivateUser
	json.Unmarshal(userJson, &privateUser)
	return privateUser
}
