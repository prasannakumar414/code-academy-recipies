-- +goose Up
-- +goose StatementBegin
create table if not exists recipies (
    id INT PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT,
    instructions TEXT,
    prep_time TEXT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS recipies;
-- +goose StatementEnd
