-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS offsets (
    id integer PRIMARY KEY,    
    offset bigint NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS offsets;
-- +goose StatementEnd
