-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS stocks (
    warehouse_id bigint PRIMARY KEY,
    sku integer NOT NULL,
    count bigint NOT NULL,
    reserved boolean NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS stocks;
-- +goose StatementEnd
