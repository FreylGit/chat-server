-- +goose Up
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL
);

CREATE TABLE chats (
    id SERIAL PRIMARY KEY
);

CREATE TABLE chat_users (
    chat_id INT REFERENCES chats(id),
    user_id INT REFERENCES users(id),
    PRIMARY KEY(chat_id, user_id)
);

CREATE TABLE messages (
    id SERIAL PRIMARY KEY,
    chat_id INT REFERENCES chats(id),
    user_id INT REFERENCES users(id),
    text TEXT NOT NULL,
    timestamp TIMESTAMP NOT NULL
);

-- +goose Down
drop table if exists messages;
drop table if exists chat_users;
drop table if exists chats;
drop table if exists users;
