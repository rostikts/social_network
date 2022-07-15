package service

import (
	"database/sql"
	"errors"
	"github.com/rostikts/social_network/domain/model"
	"github.com/rostikts/social_network/pkg/user"
	"github.com/rostikts/social_network/pkg/user/http_errors"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	repository user.Repository
}

func NewUserService(repo user.Repository) user.Service {
	return userService{repository: repo}
}

func (s userService) RegisterUser(user model.User) (model.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return model.User{}, err
	}
	user.Password = string(hashedPassword)
	res, err := s.repository.Create(user)
	if err != nil {
		return model.User{}, err
	}
	return res, err
}

func (s userService) UpdateUserData(user model.User) (model.User, error) {
	res, err := s.repository.Update(user)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.User{}, http_errors.ErrUserNotFound
		} else {
			return model.User{}, http_errors.GeneralError
		}
	}
	return res, nil
}

func (s userService) ChangePassword(user model.User) error {
	if err := s.repository.UpdatePassword(user); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return http_errors.ErrUserNotFound
		} else {
			return http_errors.GeneralError
		}
	}
	return nil
}

func (s userService) GetUserByID(id uint) (model.User, error) {
	res, err := s.repository.FindByID(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.User{}, http_errors.ErrUserNotFound
		} else {
			return model.User{}, http_errors.GeneralError
		}

	}
	return res, nil
}

func (s userService) GetAllUsers() ([]model.User, error) {
	res, err := s.repository.FindAll()
	if err != nil {
		return []model.User{}, http_errors.GeneralError
	}
	return res, nil
}
