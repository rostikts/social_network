package testutils

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"testing"
)

func NewDbMock(t *testing.T) (*sqlx.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("error occurred during mock init")
	}
	sqlxMock := sqlx.NewDb(db, "mock_db")
	return sqlxMock, mock, err
}
