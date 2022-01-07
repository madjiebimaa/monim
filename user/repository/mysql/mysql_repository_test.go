package mysql_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/madjiebimaa/monim/domain"
	repository "github.com/madjiebimaa/monim/user/repository/mysql"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestGetByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	userID := uuid.NewString()
	rows := sqlmock.NewRows([]string{"id", "name", "email", "password", "created_at"}).
		AddRow(userID, "Adjie", "adjie@gmail.com", "adjie", time.Now())

	query := "SELECT id, name, email, password, created_at FROM users WHERE id = \\?"

	t.Run("success get user", func(t *testing.T) {
		mock.ExpectQuery(query).WithArgs(userID).WillReturnRows(rows)

		m := repository.NewMysqlAuthorRepository(db)

		user, err := m.GetByID(context.TODO(), userID)
		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, userID, user.ID)
	})

	t.Run("fail because invalid argument", func(t *testing.T) {
		userID = "t"
		mock.ExpectQuery(query).WithArgs(userID).WillReturnError(domain.ErrBadParamInput)

		m := repository.NewMysqlAuthorRepository(db)

		user, err := m.GetByID(context.TODO(), userID)
		assert.Error(t, err)
		assert.Equal(t, domain.User{}, user)
	})
}
