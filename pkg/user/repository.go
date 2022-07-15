package user

import "github.com/rostikts/social_network/domain/model"

type Repository interface {
	Create(user model.User) (model.User, error)
	Update(user model.User) (model.User, error)
	UpdatePassword(user model.User) error
	FindByID(id uint) (model.User, error)
	FindAll() ([]model.User, error)
}
