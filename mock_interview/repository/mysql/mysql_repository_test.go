package mysql_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/madjiebimaa/monim/domain"
	repository "github.com/madjiebimaa/monim/mock_interview/repository/mysql"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestGetByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	query := "SELECT id, language, programming_language, meet_at, created_at, updated_at from mock_interview WHERE id = \\?"

	miID := uuid.NewString()
	now := time.Now()
	rows := sqlmock.NewRows([]string{"id", "language", "programming_language", "meet_at", "created_at", "updated_at"}).
		AddRow(miID, "english", "go", now, now, now)

	t.Run("success get mock interview", func(t *testing.T) {
		mock.ExpectQuery(query).WithArgs(miID).WillReturnRows(rows)

		m := repository.NewMysqlMockInterviewRepository(db)

		mi, err := m.GetByID(context.TODO(), miID)
		assert.NoError(t, err)
		assert.NotNil(t, mi)
		assert.Equal(t, miID, mi.ID)
	})

	t.Run("fail because invalid argument", func(t *testing.T) {
		miID = "t"
		mock.ExpectQuery(query).WithArgs(miID).WillReturnError(domain.ErrBadParamInput)

		m := repository.NewMysqlMockInterviewRepository(db)

		mi, err := m.GetByID(context.TODO(), miID)
		assert.Error(t, err)
		assert.Equal(t, domain.MockInterview{}, mi)
	})
}

func TestUpdate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	query := "UPDATE mock_interviews SET meet_at = \\?, updated_at = \\? WHERE id = \\?"

	now := time.Now()
	mi := domain.MockInterview{
		ID:        "test",
		MeetAt:    now,
		CreatedAt: now,
		UpdatedAt: now,
	}

	t.Run("success update mock interview", func(t *testing.T) {
		mock.ExpectExec(query).WithArgs(mi.MeetAt, mi.UpdatedAt, mi.ID).WillReturnResult(sqlmock.NewResult(0, 1))

		m := repository.NewMysqlMockInterviewRepository(db)

		err = m.Update(context.TODO(), &mi)
		assert.NoError(t, err)
	})

	t.Run("fail because invalid argument", func(t *testing.T) {
		mock.ExpectExec(query).WithArgs(mi.MeetAt, mi.UpdatedAt, mi.Language).WillReturnError(domain.ErrBadParamInput)

		m := repository.NewMysqlMockInterviewRepository(db)

		err = m.Update(context.TODO(), &mi)
		assert.Error(t, err)
	})

}

func TestDelete(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	query := "DELETE FROM mock_interviews WHERE id = \\?"
	miID := uuid.NewString()

	t.Run("success delete mock interview", func(t *testing.T) {
		mock.ExpectExec(query).WithArgs(miID).WillReturnResult(sqlmock.NewResult(0, 1))

		m := repository.NewMysqlMockInterviewRepository(db)

		err = m.Delete(context.TODO(), miID)
		assert.NoError(t, err)
	})

	t.Run("fail because invalid argument", func(t *testing.T) {
		miID = "t"
		mock.ExpectExec(query).WithArgs(miID).WillReturnError(domain.ErrBadParamInput)

		m := repository.NewMysqlMockInterviewRepository(db)

		err = m.Delete(context.TODO(), miID)
		assert.Error(t, err)
	})

}
