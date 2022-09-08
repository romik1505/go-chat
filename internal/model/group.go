package model

import "database/sql"

type Group struct {
	ID   sql.NullString `db:"id"`
	Name sql.NullString `db:"name"`
	Img  sql.NullString `db:"img"`

	UserID   sql.NullString `db:"user_id"` // 1-n
	UserName sql.NullString `db:"username"`
	UserImg  sql.NullString `db:"user_img"`
}

type GroupRequest struct {
	ID      sql.NullString `db:"id"`
	UserID  sql.NullString `db:"user_id"`
	GroupID sql.NullString `db:"group_id"`
	Type    sql.NullString `db:"type"` // TODO [person_join...]
	Date    sql.NullTime   `db:"date"`
}
