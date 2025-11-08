-- +goose Up
CREATE TYPE income_adjustment_category AS ENUM ('overtime', 'deduction', 'other');

CREATE TABLE income_adjustments (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    category income_adjustment_category NOT NULL,
    amount DECIMAL(10, 2) NOT NULL,
    reason VARCHAR(255),
    adjustment_date DATE NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

COMMENT ON COLUMN income_adjustments.adjustment_date IS '例: 2025年11月1日';

-- +goose Down
DROP TABLE IF EXISTS income_adjustments;

DROP TYPE IF EXISTS income_adjustment_category;