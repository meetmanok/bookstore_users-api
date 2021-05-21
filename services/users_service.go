package services

import (
	"github.com/meetmanok/bookstore_users-api/domain/users"
	"github.com/meetmanok/bookstore_users-api/utils/crypto_utils"
	"github.com/meetmanok/bookstore_users-api/utils/date_utils"
	"github.com/meetmanok/bookstore_users-api/utils/errors"
)

var(
	UsersService usersService = usersService{}
)

type usersService struct {
}

type userServiceInterface interface {
	GetUser(int64) (*users.User, *errors.RestErr)
	CreateUser(users.User) (*users.User, *errors.RestErr)
	UpdateUser(bool, users.User) (*users.User, *errors.RestErr)
	DeleteUser(int64) *errors.RestErr
	SearchUser(string) (users.Users, *errors.RestErr)
	LoginUser(users.LoginRequest) (users.Users, *errors.RestErr)
}

func (us *usersService) GetUser(userId int64) (*users.User, *errors.RestErr) {
	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func (us *usersService) CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	user.Status = users.StatusActive
	user.DateCreated = date_utils.GetNowDBFormat()
	user.Password = crypto_utils.GetMD5(user.Password)
	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func (us *usersService) UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr) {
	current, err := us.GetUser(user.Id)
	if err != nil {
		return nil, err
	}

	if isPartial {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}
		if user.Email != "" {
			current.Email = user.Email
		}
	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
	}
	current.DateCreated = date_utils.GetNowDBFormat()

	if err := current.Update(); err != nil {
		return nil, err
	}

	return current, nil
}

func (us *usersService) DeleteUser(userId int64) *errors.RestErr {
	result := &users.User{Id: userId}
	if err := result.Delete(); err != nil {
		return err
	}
	return nil
}

func (us *usersService) SearchUser(status string) (users.Users, *errors.RestErr){
	dao := &users.User{}
	return dao.FindByStatus(status)
}

func (us *usersService) LoginUser(loginreq users.LoginRequest) (*users.User, *errors.RestErr){
	dao := &users.User{
		Email:loginreq.Email,
		Password: crypto_utils.GetMD5(loginreq.Password),
	}

	if err := dao.FindByEmailAndPassword(); err != nil{
		return nil, err
	}
	return dao, nil
}