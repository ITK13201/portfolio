-- +goose Up
-- +goose StatementBegin
CREATE TABLE images
(
    id         INTEGER AUTO_INCREMENT NOT NULL PRIMARY KEY,
    path       VARCHAR(256)           NOT NULL,
    created_at DATETIME               NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME               NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS images;
-- +goose StatementEnd
