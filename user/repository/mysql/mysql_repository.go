package mysql

import (
	"context"
	"database/sql"

	"github.com/madjiebimaa/monim/domain"
)

type mysqlUserRepo struct {
	DB *sql.DB
}

func NewMysqlAuthorRepository(DB *sql.DB) domain.UserRepository {
	return &mysqlUserRepo{
		DB,
	}
}

func (m *mysqlUserRepo) GetByID(ctx context.Context, id string) (domain.User, error) {
	query := `SELECT id, name, email, password, created_at FROM users WHERE id = ?`

	row, err := m.DB.QueryContext(ctx, query, id)
	if err != nil {
		return domain.User{}, err
	}

	var user domain.User
	if row.Next() {
		if err := row.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.Password,
			&user.CreatedAt,
		); err != nil {
			return domain.User{}, err
		}
	}

	return user, nil
}
