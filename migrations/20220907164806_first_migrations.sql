-- +goose Up
-- +goose StatementBegin

CREATE TABLE users (
    id TEXT DEFAULT gen_random_uuid() PRIMARY KEY NOT NULL,
    username VARCHAR(100) NOT NULL,
    img VARCHAR(100),
    password TEXT NOT NULL
);

CREATE TABLE sessions (
    id TEXT DEFAULT gen_random_uuid() PRIMARY KEY NOT NULL,
    user_id TEXT NOT NULL,
    token TEXT NOT NULL,
    expires_in INTEGER NOT NULL,
    device_id TEXT NOT NULL
);

CREATE TABLE groups (
    id TEXT DEFAULT gen_random_uuid() PRIMARY KEY NOT NULL,
    name TEXT NOT NULL,
    img TEXT NOT NULL,
    created_at timestamptz NOT NULL DEFAULT now()
);

CREATE TABLE user_groups(
    user_id TEXT REFERENCES users(id),
    group_id TEXT REFERENCES groups(id) ON DELETE CASCADE
);

CREATE TABLE messages(
    id TEXT DEFAULT gen_random_uuid() PRIMARY KEY NOT NULL,
    sender_id TEXT NOT NULL REFERENCES users(id),
    receiver_id TEXT NOT NULL,
    receiver_type TEXT NOT NULL,
    text TEXT NOT NULL,
    created_at timestamptz NOT NULL DEFAULT now()
);

CREATE TABLE friends(
    user_id TEXT REFERENCES users(id),
    user_friend TEXT REFERENCES users(id),
    status TEXT NOT NULL,
    CONSTRAINT PK_T1 PRIMARY KEY(user_id, user_friend)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE user_groups;
DROP TABLE messages;
DROP TABLE users;
DROP TABLE sessions;
DROP TABLE groups;
-- +goose StatementEnd
