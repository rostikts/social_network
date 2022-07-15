package mocks

import "github.com/rostikts/social_network/domain/model"

type UserRepoMock struct {
	CreateFunc         func(user model.User) (model.User, error)
	UpdateFunc         func(user model.User) (model.User, error)
	UpdatePasswordFunc func(user model.User) error
	FindByIDFunc       func(id uint) (model.User, error)
	FindAllFunc        func() ([]model.User, error)
}

func (m UserRepoMock) Create(user model.User) (model.User, error) {
	return m.CreateFunc(user)
}
func (m UserRepoMock) Update(user model.User) (model.User, error) {
	return m.UpdateFunc(user)
}
func (m UserRepoMock) UpdatePassword(user model.User) error {
	return m.UpdatePasswordFunc(user)
}
func (m UserRepoMock) FindByID(id uint) (model.User, error) {
	return m.FindByIDFunc(id)
}
func (m UserRepoMock) FindAll() ([]model.User, error) {
	return m.FindAllFunc()
}
