-- +goose Up
-- +goose StatementBegin
ALTER TYPE item_status RENAME VALUE ' completed' TO 'completed';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TYPE item_status RENAME VALUE 'completed' TO ' completed';
-- +goose StatementEnd
