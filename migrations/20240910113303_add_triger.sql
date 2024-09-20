-- +goose Up

CREATE OR REPLACE FUNCTION check_user_chat()
    RETURNS TRIGGER AS $$
BEGIN
    IF NOT EXISTS (
        SELECT 1
        FROM chat_user
        WHERE user_id = NEW.user_id AND chat_id = NEW.chat_id
    ) THEN
        RAISE EXCEPTION 'Запись с user_id % и chat_id % не найдена в таблице chat_user', NEW.user_id, NEW.chat_id;
END IF;

RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER check_user_chat_before_insert
    BEFORE INSERT ON message
    FOR EACH ROW
    EXECUTE FUNCTION check_user_chat();

-- +goose Down
