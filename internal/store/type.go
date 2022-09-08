package store

import (
	"database/sql"
	"time"
)

func NewNullString(v string) sql.NullString {
	return sql.NullString{v, true}
}

func NewNullTime(t time.Time) sql.NullTime {
	return sql.NullTime{t, true}
}
