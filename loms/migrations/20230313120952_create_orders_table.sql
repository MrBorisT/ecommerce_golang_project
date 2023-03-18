-- +goose Up
-- +goose StatementBegin
CREATE TYPE status AS ENUM ('new', 'awaiting payment', 'failed', 'payed', 'cancelled');

CREATE TABLE IF NOT EXISTS orders (
    id bigint GENERATED ALWAYS AS IDENTITY ,
    status status NOT NULL,
    user_id bigint NOT NULL,
    PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS order_items (
    order_id bigint NOT NULL,
    sku integer NOT NULL,
    count integer NOT NULL,
    CONSTRAINT fk_order_id
        FOREIGN KEY(order_id) 
        REFERENCES orders(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS order_items;
DROP TABLE IF EXISTS orders;
DROP TYPE status;
-- +goose StatementEnd
