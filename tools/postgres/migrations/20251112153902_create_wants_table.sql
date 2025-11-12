-- +goose Up
-- +goose StatementBegin
CREATE TABLE wants (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users (id),
    name VARCHAR(255) NOT NULL, -- 欲しいものの名前
    target_amount NUMERIC(10, 2) NOT NULL, -- 目標金額
    target_date DATE NOT NULL, -- 目標購入日
    purchased BOOLEAN DEFAULT FALSE, -- 購入済みフラグ
    purchased_at DATE NULL, -- 購入日
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- 作成日時
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP -- 更新日時
);

CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_wants_updated_at
BEFORE UPDATE ON wants
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS wants;
-- +goose StatementEnd