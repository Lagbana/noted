-- +goose Up
-- +goose StatementBegin
CREATE TYPE list_status AS ENUM ('active', ' completed', 'cancelled');

CREATE TABLE todo_list (
    id SERIAL PRIMARY KEY,
    status list_status NOT NULL DEFAULT 'active',
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS todo_list;
-- +goose StatementEnd
