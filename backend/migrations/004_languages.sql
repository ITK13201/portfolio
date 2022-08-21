-- +goose Up
-- +goose StatementBegin
CREATE TABLE languages
(
    id         INTEGER AUTO_INCREMENT NOT NULL PRIMARY KEY,
    name       VARCHAR(32)            NOT NULL UNIQUE,
    image_id   INTEGER                NOT NULL,
    created_at DATETIME               NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME               NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_languages_image_id
        FOREIGN KEY (image_id)
            REFERENCES images (id)
            ON DELETE RESTRICT ON UPDATE RESTRICT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS languages;
-- +goose StatementEnd
