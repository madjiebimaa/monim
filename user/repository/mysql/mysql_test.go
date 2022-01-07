package mysql_test

import (
	"context"
	"testing"
	"time"

	repository "github.com/madjiebimaa/monim/user/repository/mysql"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestGetByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	rows := sqlmock.NewRows([]string{"id", "name", "email", "password", "created_at"}).
		AddRow("adjie", "Adjie", "adjie@gmail.com", "adjie", time.Now())

	query := "SELECT id, name, email, password, created_at FROM users WHERE id = \\?"

	userID := "adjie"
	mock.ExpectQuery(query).WithArgs(userID).WillReturnRows(rows)

	m := repository.NewMysqlAuthorRepository(db)

	anUser, err := m.GetByID(context.TODO(), userID)
	assert.NoError(t, err)
	assert.NotNil(t, anUser)

}
