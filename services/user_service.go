package services

import (
	"github.com/manishsinghbisht/greenleaf-api/domain/user"
	"github.com/manishsinghbisht/utils-go/date_utils"
	"github.com/manishsinghbisht/utils-go/crypto_utils"
	"github.com/manishsinghbisht/utils-go/rest_errors"
)

var (
	UserService userServiceInterface = &userService{}
)

type userService struct{}

type userServiceInterface interface {
	GetUser(int64) (*user.User, rest_errors.RestErr)
	CreateUser(user.User) (*user.User, rest_errors.RestErr)
	UpdateUser(bool, user.User) (*user.User, rest_errors.RestErr)
	DeleteUser(int64) rest_errors.RestErr
	SearchUser(string) (user.Users, rest_errors.RestErr)
	LoginUser(user.LoginRequest) (*user.User, rest_errors.RestErr)
}

func (s *userService) GetUser(userId int64) (*user.User, rest_errors.RestErr) {
	dao := &user.User{Id: userId}
	if err := dao.Get(); err != nil {
		return nil, err
	}
	return dao, nil
}

func (s *userService) CreateUser(user_dto user.User) (*user.User, rest_errors.RestErr) {
	if err := user_dto.Validate(); err != nil {
		return nil, err
	}

	user_dto.Status = user.StatusActive
	user_dto.DateCreated = date_utils.GetNowDBFormat()
	user_dto.Password = crypto_utils.GetMd5(user_dto.Password)
	if err := user_dto.Save(); err != nil {
		return nil, err
	}
	return &user_dto, nil
}

func (s *userService) UpdateUser(isPartial bool, user_dto user.User) (*user.User, rest_errors.RestErr) {
	current := &user.User{Id: user_dto.Id}
	if err := current.Get(); err != nil {
		return nil, err
	}

	if isPartial {
		if user_dto.FirstName != "" {
			current.FirstName = user_dto.FirstName
		}

		if user_dto.LastName != "" {
			current.LastName = user_dto.LastName
		}

		if user_dto.Email != "" {
			current.Email = user_dto.Email
		}
	} else {
		current.FirstName = user_dto.FirstName
		current.LastName = user_dto.LastName
		current.Email = user_dto.Email
	}

	if err := current.Update(); err != nil {
		return nil, err
	}
	return current, nil
}

func (s *userService) DeleteUser(userId int64) rest_errors.RestErr {
	dao := &user.User{Id: userId}
	return dao.Delete()
}

func (s *userService) SearchUser(status string) (user.Users, rest_errors.RestErr) {
	dao := &user.User{}
	return dao.FindByStatus(status)
}

func (s *userService) LoginUser(request user.LoginRequest) (*user.User, rest_errors.RestErr) {
	dao := &user.User{
		Email:    request.Email,
		Password: crypto_utils.GetMd5(request.Password),
	}
	if err := dao.FindByEmailAndPassword(); err != nil {
		return nil, err
	}
	return dao, nil
}
