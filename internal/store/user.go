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

func (s Storage) GetUser(ctx context.Context, user model.User) (model.User, error) {
	query := s.Builder().Select("*").From("users").Where(sq.Eq{"username": user.Username, "password": user.Password})
	err := query.QueryRow().Scan(&user)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (s Storage) UserList(ctx context.Context) ([]model.User, error) {
	q := s.Builder().Select("*").From("users").GroupBy("id")

	// rows, err := query.Query()
	users := make([]model.User, 0)

	query, _, _ := q.ToSql()
	rows, err := s.Queryx(query)

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
