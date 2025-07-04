-- +goose Up
-- +goose StatementBegin
create table if not exists recipies (
    ID TEXT PRIMARY KEY,
    Name TEXT NOT NULL,
    Descriptions TEXT,
    Instructions TEXT,
    PrepTime TEXT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS recipies;
-- +goose StatementEnd
