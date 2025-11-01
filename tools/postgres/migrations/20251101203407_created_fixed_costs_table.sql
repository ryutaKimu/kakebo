-- +goose Up
-- +goose StatementBegin
CREATE TABLE fixed_costs (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    name VARCHAR(255) NOT NULL,
    amount NUMERIC(10,2) NOT NULL,
    payment_date SMALLINT,
    memo TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS fixed_costs;
-- +goose StatementEnd
