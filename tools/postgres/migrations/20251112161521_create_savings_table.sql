-- +goose Up
-- +goose StatementBegin
CREATE TABLE savings (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users (id),
    amount DECIMAL(10, 2) NOT NULL,
    comment VARCHAR(255),
    saved_at DATE NOT NULL
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS savings;
-- +goose StatementEnd