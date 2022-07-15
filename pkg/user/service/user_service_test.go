package service

import (
	"database/sql"
	"errors"
	"github.com/rostikts/social_network/domain/model"
	"github.com/rostikts/social_network/pkg/user/http_errors"
	"github.com/rostikts/social_network/pkg/user/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

var defaultUser = model.User{
	ID:        1,
	UserName:  "test",
	FirstName: "test1",
	LastName:  "last1",
	Email:     "test@last.com",
}

func TestUserService_GetUserByID(t *testing.T) {
	type mock struct {
		user model.User
		err  error
	}
	type result struct {
		user model.User
		err  error
	}
	tests := []struct {
		name   string
		userId uint
		result result
		mock   mock
	}{
		{
			name:   "get single user by id",
			userId: 1,
			result: result{
				user: defaultUser,
				err:  nil,
			},
			mock: mock{
				user: defaultUser,
				err:  nil,
			},
		},
		{
			name:   "get non existing user",
			userId: 1,
			result: result{
				user: model.User{},
				err:  http_errors.ErrUserNotFound,
			},
			mock: mock{
				user: model.User{},
				err:  sql.ErrNoRows,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			service := NewUserService(mocks.UserRepoMock{FindByIDFunc: func(id uint) (model.User, error) {
				return tc.mock.user, tc.mock.err
			}})
			res, err := service.GetUserByID(tc.userId)
			assert.Equal(t, tc.result.user, res)
			assert.ErrorIs(t, tc.result.err, err)
		})
	}
}

func TestUserService_GetAllUsers(t *testing.T) {
	type result struct {
		users []model.User
		err   error
	}
	tests := []struct {
		name    string
		result  result
		mockErr error
	}{
		{
			name: "get single user in list",
			result: result{
				users: []model.User{defaultUser},
				err:   nil,
			},
		},
		{
			name: "get several users in list",
			result: result{
				users: []model.User{defaultUser, defaultUser},
				err:   nil,
			},
		},
		{
			name: "get users with error",
			result: result{
				users: []model.User{},
				err:   http_errors.GeneralError,
			},
			mockErr: errors.New("some error"),
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			service := NewUserService(mocks.UserRepoMock{FindAllFunc: func() ([]model.User, error) {
				return tc.result.users, tc.mockErr
			}})
			res, err := service.GetAllUsers()
			assert.Equal(t, tc.result.users, res)
			assert.ErrorIs(t, tc.result.err, err)
		})
	}
}

func TestUserService_UpdateUserData(t *testing.T) {
	type mock struct {
		user model.User
		err  error
	}
	type result struct {
		user model.User
		err  error
	}
	tests := []struct {
		name    string
		userArg model.User
		result  result
		mock    mock
	}{
		{
			name:    "update user",
			userArg: defaultUser,
			result: result{
				user: defaultUser,
				err:  nil,
			},
			mock: mock{
				user: defaultUser,
				err:  nil,
			},
		},
		{
			name:    "updated not existing user",
			userArg: defaultUser,
			result: result{
				user: model.User{},
				err:  http_errors.ErrUserNotFound,
			},
			mock: mock{
				user: model.User{},
				err:  sql.ErrNoRows,
			},
		},
		{
			name:    "get users with error",
			userArg: defaultUser,
			result: result{
				user: model.User{},
				err:  http_errors.GeneralError,
			},
			mock: mock{
				user: model.User{},
				err:  errors.New("some error"),
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			service := NewUserService(mocks.UserRepoMock{UpdateFunc: func(user model.User) (model.User, error) {
				return tc.mock.user, tc.mock.err
			}})
			res, err := service.UpdateUserData(defaultUser)
			assert.Equal(t, tc.result.user, res)
			assert.ErrorIs(t, tc.result.err, err)
		})
	}
}

func TestUserService_ChangePassword(t *testing.T) {
	tests := []struct {
		name    string
		userArg model.User
		result  error
		mock    error
	}{
		{
			name:    "update password",
			userArg: defaultUser,
			result:  nil,
			mock:    nil,
		},
		{
			name:    "update password of non existing user",
			userArg: defaultUser,
			result:  http_errors.ErrUserNotFound,
			mock:    sql.ErrNoRows,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			service := NewUserService(mocks.UserRepoMock{UpdatePasswordFunc: func(user model.User) error {
				return tc.mock
			}})
			err := service.ChangePassword(tc.userArg)
			assert.ErrorIs(t, err, tc.result)
		})
	}
}

func TestUserService_RegisterUser(t *testing.T) {
	type mock struct {
		user model.User
		err  error
	}
	type result struct {
		user model.User
		err  error
	}
	tests := []struct {
		name    string
		userArg model.User
		result  result
		mock    mock
	}{
		{
			name: "update user",
			userArg: model.User{
				ID:        1,
				UserName:  "test",
				FirstName: "test1",
				LastName:  "last1",
				Email:     "test@last.com",
				Password:  "1234215",
			},
			result: result{
				user: defaultUser,
				err:  nil,
			},
			mock: mock{
				defaultUser,
				nil,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			service := NewUserService(mocks.UserRepoMock{CreateFunc: func(user model.User) (model.User, error) {
				return user, nil
			}})
			res, err := service.RegisterUser(tc.userArg)
			assert.NotEqual(t, tc.userArg.Password, res.Password)
			assert.ErrorIs(t, err, tc.result.err)
		})
	}
}
