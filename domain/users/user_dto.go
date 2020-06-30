package users

import (
	// "database/sql"
	"strings"

	"github.com/manishsinghbisht/utils-go/rest_errors"
	"gopkg.in/guregu/null.v4"
)

const (
	StatusActive = "active"
)

type User struct {
	Id          int64  `json:"Id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Gender      null.String `json:"gender"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	Password    string `json:"password"`
	Mobile     int64 `json:"mobile"`
	MobilePin      int64 `json:"mobile_pin"`
	PostalPin      int64 `json:"postal_pin"`
}

type Users []User

func (user *User) Validate() rest_errors.RestErr {
	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)

	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return rest_errors.NewBadRequestError("invalid email address")
	}

	user.Password = strings.TrimSpace(user.Password)
	if user.Password == "" {
		return rest_errors.NewBadRequestError("invalid password")
	}
	return nil
}
