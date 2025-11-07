-- +goose Up
CREATE TYPE income_adjustment_category AS ENUM ('overtime', 'deduction', 'other');

CREATE TABLE income_adjustments (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    category income_adjustment_category NOT NULL,
    amount DECIMAL(10, 2) NOT NULL,
    reason VARCHAR(255),
    adjustment_month CHAR(7) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

COMMENT ON COLUMN income_adjustments.adjustment_month IS '例: 11';

-- +goose Down
DROP TABLE IF EXISTS income_adjustments;

DROP TYPE IF EXISTS income_adjustment_category;