-- +goose Up
-- +goose StatementBegin
CREATE TYPE item_status AS ENUM ('active', ' completed', 'cancelled');

CREATE TABLE todo_items (
    id SERIAL PRIMARY KEY,
    todo_list_id SERIAL NOT NULL,
    status item_status NOT NULL DEFAULT 'active',
    todo text NOT NULL DEFAULT '',
    comment text NOT NULL DEFAULT '',
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz,
    CONSTRAINT fk_todo_list FOREIGN KEY (todo_list_id) references todo_list (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS todo_items;
-- +goose StatementEnd
