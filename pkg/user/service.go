package user

import "github.com/rostikts/social_network/domain/model"

type Service interface {
	RegisterUser(user model.User) (model.User, error)
	UpdateUserData(user model.User) (model.User, error)
	ChangePassword(user model.User) error
	GetUserByID(id uint) (model.User, error)
	GetAllUsers() ([]model.User, error)
}
