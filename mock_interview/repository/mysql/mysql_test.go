package mysql_test

import (
	"context"
	"testing"
	"time"

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

	now := time.Now()
	rows := sqlmock.NewRows([]string{"id", "language", "programming_language", "meet_at", "created_at", "updated_at"}).
		AddRow("test", "english", "go", now, now, now)

	query := "SELECT id, language, programming_language, meet_at, created_at, updated_at from mock_interview WHERE id = \\?"

	miID := "test"
	mock.ExpectQuery(query).WithArgs(miID).WillReturnRows(rows)

	m := repository.NewMysqlMockInterviewRepository(db)

	anUser, err := m.GetByID(context.TODO(), miID)
	assert.NoError(t, err)
	assert.NotNil(t, anUser)
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

	mock.ExpectExec(query).WithArgs(mi.MeetAt, mi.UpdatedAt, mi.ID).WillReturnResult(sqlmock.NewResult(0, 1))

	m := repository.NewMysqlMockInterviewRepository(db)

	err = m.Update(context.TODO(), &mi)
	assert.NoError(t, err)
}

func TestDelete(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	query := "DELETE FROM mock_interviews WHERE id = \\?"

	miID := "test"
	mock.ExpectExec(query).WithArgs(miID).WillReturnResult(sqlmock.NewResult(0, 1))

	m := repository.NewMysqlMockInterviewRepository(db)

	err = m.Delete(context.TODO(), miID)
	assert.NoError(t, err)
}
