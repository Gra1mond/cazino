-- +goose Up
ALTER TABLE users
    ADD COLUMN email TEXT NOT NULL,
    ADD COLUMN password_hash TEXT NOT NULL;

CREATE UNIQUE INDEX user_email_index
    ON users(lower(email));

-- +goose Down
DROP INDEX user_email_index;

ALTER TABLE
    DROP COLUMN email,
    DROP COLUMN password_hash;
