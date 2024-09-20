-- +goose Up
CREATE TABLE chat (
  id serial PRIMARY KEY
);

CREATE TABLE message (
    id serial PRIMARY KEY ,
    user_id bigint not null,
    chat_id bigint not null references chat(id) ON DELETE CASCADE,
    text text not null,
    create_at timestamp not null default now()
);

CREATE TABLE chat_user(
    user_id bigint not null ,
    chat_id bigint not null references chat(id) ON DELETE CASCADE
);

-- +goose Down
drop table chat_user;
drop table message;
drop table chat;
