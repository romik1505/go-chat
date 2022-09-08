package store

import (
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Storage struct {
	*sqlx.DB
}

func (s Storage) Builder() sq.StatementBuilderType {
	return sq.StatementBuilder.PlaceholderFormat(sq.Dollar).RunWith(s.DB)
}

func NewStorage() Storage {
	conStr := "postgresql://postgres:1505@localhost:5432/chat?sslmode=disable&timezone=UTC"
	db, err := sqlx.Open("postgres", conStr)
	if err != nil {
		log.Fatalln("database err: ", err.Error())
	}
	return Storage{
		DB: db,
	}
}
