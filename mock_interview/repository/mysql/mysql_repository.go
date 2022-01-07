package mysql

import (
	"context"
	"database/sql"

	"github.com/madjiebimaa/monim/domain"
)

type mysqlMockInterviewRepo struct {
	DB *sql.DB
}

func NewMysqlMockInterviewRepository(DB *sql.DB) domain.MockInterviewRepository {
	return &mysqlMockInterviewRepo{
		DB,
	}
}

func (m *mysqlMockInterviewRepo) GetByID(ctx context.Context, id string) (domain.MockInterview, error) {
	query := `SELECT id, language, programming_language, meet_at, created_at, updated_at from mock_interview WHERE id = ?`
	row, err := m.DB.QueryContext(ctx, query, id)
	if err != nil {
		return domain.MockInterview{}, nil
	}

	var mi domain.MockInterview
	if row.Next() {
		row.Scan(
			&mi.ID,
			&mi.Language,
			&mi.ProgrammingLanguage,
			&mi.MeetAt,
			&mi.CreatedAt,
			&mi.UpdatedAt,
		)
	}

	return mi, nil
}

func (m *mysqlMockInterviewRepo) Update(ctx context.Context, mi *domain.MockInterview) error {
	query := `UPDATE mock_interviews SET meet_at = ?, updated_at = ? WHERE id = ?`
	if _, err := m.DB.ExecContext(ctx, query, mi.MeetAt, mi.UpdatedAt, mi.ID); err != nil {
		return err
	}

	return nil
}

func (m *mysqlMockInterviewRepo) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM mock_interviews WHERE id = ?`
	if _, err := m.DB.ExecContext(ctx, query, id); err != nil {
		return err
	}

	return nil
}
