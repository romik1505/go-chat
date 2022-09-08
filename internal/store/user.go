package store

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/romik1505/chat/internal/model"
)

func (s Storage) CreateUser(ctx context.Context, user model.User) (model.User, error) {
	err := s.Builder().
		Insert("users").
		Columns("username", "password").
		Values(user.Username, user.Password).
		Suffix("RETURNING id").QueryRow().Scan(&user.ID)
	return user, err
}

func applyUserFilters(query sq.SelectBuilder, user model.User) sq.SelectBuilder {
	if user.ID.Valid {
		query = query.Where(sq.Eq{"id": user.ID})
	}
	if user.Username.Valid {
		query = query.Where(sq.Eq{"username": user.Username})
	}
	if user.Password.Valid {
		query = query.Where(sq.Eq{"password": user.Password})
	}
	return query
}

func (s Storage) GetUser(ctx context.Context, user model.User) (model.User, error) {
	query := s.Builder().Select("*").From("users")
	query = applyUserFilters(query, user)

	sql, vars := query.MustSql()
	err := s.QueryRowx(sql, vars...).StructScan(&user)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (s Storage) UserList(ctx context.Context) ([]model.User, error) {
	q := s.Builder().Select("*").From("users").OrderBy("id")

	users := make([]model.User, 0)

	query, _ := q.MustSql()
	rows, err := s.Queryx(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var user model.User
		err := rows.StructScan(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, err
}
