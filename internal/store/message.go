package store

import (
	"context"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/romik1505/chat/internal/model"
)

type MessageSearchOpts struct {
	ChatID   string
	ChatType string
	UserID   string
}

func applyMessageFilter(q sq.SelectBuilder, opts MessageSearchOpts) sq.SelectBuilder {
	q = q.Where(sq.Eq{"receiver_type": opts.ChatType})

	if opts.ChatType == "group" {
		q = q.Where(sq.Eq{"receiver_id": opts.ChatID})
	}
	if opts.ChatType == "person" {
		q = q.Where(sq.Eq{"receiver_id": []string{opts.ChatID, opts.UserID}})
		q = q.Where(sq.Eq{"sender_id": []string{opts.ChatID, opts.UserID}})
	}

	return q
}

func (s Storage) MessageList(ctx context.Context, opts MessageSearchOpts) ([]model.StoredMessage, error) {
	q := s.Builder().Select("*").From("messages").OrderBy("DESC created_at")

	q = applyMessageFilter(q, opts)

	sql, _, _ := q.ToSql()
	rows, err := s.Queryx(sql)
	if err != nil {
		return nil, err
	}

	res := make([]model.StoredMessage, 0)
	for rows.Next() {
		var message model.StoredMessage
		err := rows.StructScan(&message)
		if err != nil {
			return nil, err
		}

		res = append(res, message)
	}
	return res, nil
}

func (s Storage) SaveMessage(ctx context.Context, message model.StoredMessage) (model.StoredMessage, error) {
	q := s.Builder().
		Insert("messages").SetMap(map[string]interface{}{
		"text":          message.Text,
		"created_at":    message.Date,
		"sender_id":     message.SenderID,
		"receiver_id":   message.ReceiverID,
		"receiver_type": message.ReceiverType,
	}).Suffix("RETURNING id, created_at as date")

	sql, data := q.MustSql()
	log.Println(sql, " - ", data)
	err := s.QueryRowx(sql, data...).StructScan(&message)
	return message, err
}
