-- +goose Up
-- +goose StatementBegin
ALTER TABLE todo_list 
ADD COLUMN title text NOT NULL default '';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE todo_list
DROP COLUMN title;
-- +goose StatementEnd