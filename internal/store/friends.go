package store

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/romik1505/chat/internal/model"
)

type FriendStatus string

const (
	FriendStatusUndefined = FriendStatus("undefined")
	FriendStatusRequested = FriendStatus("requested")
	FriendStatusAccepted  = FriendStatus("accepted")
)

func (s Storage) FriendList(ctx context.Context, user_id string, status FriendStatus) ([]model.User, error) {
	q := s.Builder().Select("u.*").
		From("friends f").
		InnerJoin("users u on f.user_friend=u.id").
		Where(sq.Eq{"user_id": user_id, "status": status})

	query, vars := q.MustSql()
	rows, err := s.Queryx(query, vars...)
	if err != nil {
		return nil, err
	}
	users := make([]model.User, 0)
	for rows.Next() {
		var user model.User
		err := rows.StructScan(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (s Storage) InsertFriendRequest(ctx context.Context, user_id, friend_id string) (int64, error) {
	return s.insertFriend(ctx, user_id, friend_id, FriendStatusRequested)
}

func (s Storage) insertFriend(ctx context.Context, user_id, friend_id string, status FriendStatus) (int64, error) {
	q := s.Builder().Insert("friends").SetMap(map[string]interface{}{
		"user_id":     user_id,
		"user_friend": friend_id,
		"status":      status,
	})

	res, err := q.ExecContext(ctx)
	if err != nil {
		return 0, err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return count, err
}

func (s Storage) AcceptFriendRequest(ctx context.Context, user_id, friend_id string) (int64, error) {
	tx, err := s.DB.Begin()
	if err != nil {
		return 0, err
	}

	qUpdate := s.Builder().Update("friends").
		Set("status", FriendStatusAccepted).
		Where(sq.Eq{"user_friend": user_id, "user_id": friend_id})

	query, vars := qUpdate.MustSql()

	res, err := tx.Exec(query, vars...)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	qInsert := s.Builder().Insert("friends").SetMap(map[string]interface{}{
		"user_id":     user_id,
		"user_friend": friend_id,
		"status":      FriendStatusAccepted,
	})
	query, vars = qInsert.MustSql()
	res, err = tx.Exec(query, vars...)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	count, err := res.RowsAffected()

	return count, err
}
