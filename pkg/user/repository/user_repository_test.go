package repository

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/rostikts/social_network/domain/model"
	"github.com/rostikts/social_network/pkg/user"
	testutils "github.com/rostikts/social_network/testing_utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"regexp"
	"testing"
)

func setup(t *testing.T) (r user.Repository, mock sqlmock.Sqlmock, teardown func()) {
	db, mock, _ := testutils.NewDbMock(t)
	r = NewUserRepository(db)
	return r, mock, func() {
		_ = db.Close()
	}
}

func TestUserRepository_FindAll(t *testing.T) {
	tests := []struct {
		name   string
		query  string
		result []model.User
	}{
		{
			name:  "single user in list",
			query: `SELECT u.id, u.username, u.first_name, u.last_name, u.email FROM users u`,
			result: []model.User{
				{
					ID:        1,
					UserName:  "rostikts",
					FirstName: "rostik",
					LastName:  "Tsyapiura",
					Email:     "rostik@test.ua",
					Password:  "",
				},
			},
		},
		{
			name:  "several users in list",
			query: `SELECT u.id, u.username, u.first_name, u.last_name, u.email FROM users u`,
			result: []model.User{
				{
					ID:        1,
					UserName:  "rostikts",
					FirstName: "rostik",
					LastName:  "Tsyapiura",
					Email:     "rostik@test.ua",
					Password:  "",
				},
				{
					ID:        2,
					UserName:  "rostikts1",
					FirstName: "rostik231",
					LastName:  "Tsyapiura",
					Email:     "rostik@tes2t.ua",
					Password:  "",
				},
			},
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			r, mock, teardown := setup(t)
			defer teardown()

			rows := mock.NewRows([]string{"id", "username", "first_name", "last_name", "email"})
			for _, u := range tc.result {
				rows.AddRow(u.ID, u.UserName, u.FirstName, u.LastName, u.Email)
			}
			mock.ExpectQuery(regexp.QuoteMeta(tc.query)).WillReturnRows(rows)
			res, err := r.FindAll()
			assert.Equal(t, tc.result, res)
			require.NoError(t, err)
		})
	}
}

func TestUserRepository_Create(t *testing.T) {
	tests := []struct {
		name   string
		query  string
		result model.User
	}{
		{
			name:  "create a user",
			query: `INSERT INTO users(username, first_name, last_name, email, password) VALUES (?, ?, ?, ?, ?) returning id`,
			result: model.User{
				ID:        1,
				UserName:  "rostikts",
				FirstName: "rostik",
				LastName:  "Tsyapiura",
				Email:     "rostik@test.ua",
				Password:  "123245",
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			r, mock, teardown := setup(t)
			defer teardown()

			rows := mock.NewRows([]string{"id"})
			rows.AddRow(tc.result.ID)
			mock.ExpectQuery(regexp.QuoteMeta(tc.query)).WillReturnRows(rows)
			res, err := r.Create(tc.result)
			assert.Equal(t, tc.result, res)
			require.NoError(t, err)
		})
	}
}

func TestUserRepository_Update(t *testing.T) {
	tests := []struct {
		name  string
		query string
		user  model.User
	}{
		{
			name:  "create a user",
			query: `UPDATE users u SET username=?, first_name=?, last_name=?, email=? WHERE u.id=?`,
			user: model.User{
				ID:        1,
				UserName:  "rostikts",
				FirstName: "rostik",
				LastName:  "Tsyapiura",
				Email:     "rostik@test.ua",
				Password:  "123245",
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			r, mock, teardown := setup(t)
			defer teardown()

			mock.ExpectExec(regexp.QuoteMeta(tc.query)).WithArgs(tc.user.UserName, tc.user.FirstName, tc.user.LastName, tc.user.Email, tc.user.ID).WillReturnResult(sqlmock.NewResult(1, 1))
			res, err := r.Update(tc.user)
			assert.Equal(t, tc.user, res)
			require.NoError(t, err)
		})
	}
}

func TestUserRepository_UpdatePassword(t *testing.T) {
	tests := []struct {
		name  string
		query string
		user  model.User
	}{
		{
			name:  "update password",
			query: `UPDATE users u SET password=? WHERE u.id=?`,
			user: model.User{
				ID:        1,
				UserName:  "rostikts",
				FirstName: "rostik",
				LastName:  "Tsyapiura",
				Email:     "rostik@test.ua",
				Password:  "123245",
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			r, mock, teardown := setup(t)
			defer teardown()

			mock.ExpectExec(regexp.QuoteMeta(tc.query)).WithArgs(tc.user.Password, tc.user.ID).WillReturnResult(sqlmock.NewResult(1, 1))
			err := r.UpdatePassword(tc.user)
			require.NoError(t, err)
		})
	}
}

func TestUserRepository_FindByID(t *testing.T) {
	tests := []struct {
		name   string
		query  string
		result model.User
	}{
		{
			name:  "get single user by id",
			query: `SELECT u.id, u.username, u.first_name, u.last_name, u.email FROM users u WHERE u.id=$1`,
			result: model.User{
				ID:        1,
				UserName:  "rostikts",
				FirstName: "rostik",
				LastName:  "Tsyapiura",
				Email:     "rostik@test.ua",
				Password:  "",
			},
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			r, mock, teardown := setup(t)
			defer teardown()

			rows := mock.NewRows([]string{"id", "username", "first_name", "last_name", "email"})
			rows.AddRow(tc.result.ID, tc.result.UserName, tc.result.FirstName, tc.result.LastName, tc.result.Email)
			mock.ExpectQuery(regexp.QuoteMeta(tc.query)).WithArgs(tc.result.ID).WillReturnRows(rows)
			res, err := r.FindByID(tc.result.ID)
			assert.Equal(t, tc.result, res)
			require.NoError(t, err)
		})
	}
}
