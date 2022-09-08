db\:create:
	goose -dir migrations create $(NAME) sql


db\:up:
	goose -dir migrations postgres "postgresql://postgres:1505@localhost:5432/chat?sslmode=disable&timezone=UTC" up


db\:down:
	goose -dir migrations postgres "postgresql://postgres:1505@localhost:5432/chat?sslmode=disable&timezone=UTC" down