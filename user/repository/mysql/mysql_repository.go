package mysql

import (
	"context"
	"database/sql"

	"github.com/madjiebimaa/monim/domain"
)

type mysqlUserRepo struct {
	db *sql.DB
}

func NewMysqlAuthorRepository(db *sql.DB) domain.UserRepository {
	return &mysqlUserRepo{
		db,
	}
}

func (m *mysqlUserRepo) GetByID(ctx context.Context, id string) (domain.User, error) {
	query := `SELECT id, name, email, password, created_at FROM users WHERE id = ?`

	row, err := m.db.QueryContext(ctx, query, id)
	if err != nil {
		return domain.User{}, err
	}

	var user domain.User
	if row.Next() {
		row.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.Password,
			&user.CreatedAt,
		)
	}

	return user, nil
}
