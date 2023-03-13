-- +goose Up
-- +goose StatementBegin
CREATE TYPE status AS ENUM ('new', 'awaiting payment', 'failed', 'payed', 'cancelled');

CREATE TABLE IF NOT EXISTS orders (
    order_id bigint PRIMARY KEY,
    status status NOT NULL,
    user_id bigint NOT NULL    
);

CREATE TABLE IF NOT EXISTS order_items (
    order_id bigint PRIMARY KEY,
    sku integer NOT NULL,
    count integer NOT NULL    
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS order_items;
DROP TABLE IF EXISTS orders;
DROP TYPE status;
-- +goose StatementEnd
