package model

import "database/sql"

type User struct {
	ID       sql.NullString `db:"id"`
	Username sql.NullString `db:"username"`
	Img      sql.NullString `db:"img"`
	Password sql.NullString `db:"password"` // hashed password
}
