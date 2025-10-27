-- +goose Up
-- +goose StatementBegin
CREATE SEQUENCE IF NOT EXISTS users_id_seq;

ALTER TABLE users ALTER COLUMN id TYPE BIGINT;

ALTER TABLE users ALTER COLUMN id SET DEFAULT nextval('users_id_seq');

SELECT setval('users_id_seq', COALESCE((SELECT MAX(id) FROM users), 0) + 1, false);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users ALTER COLUMN id DROP DEFAULT;
DROP SEQUENCE IF EXISTS users_id_seq;
-- +goose StatementEnd
