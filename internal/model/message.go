package model

import "database/sql"

type StoredMessage struct {
	ID           sql.NullString `db:"id"`
	Text         sql.NullString `db:"text"`
	Date         sql.NullTime   `db:"date"`
	SenderID     sql.NullString `db:"sender_id"`
	ReceiverID   sql.NullString `db:"receiver_id"`
	ReceiverType sql.NullString `db:"receiver_type"`
}
