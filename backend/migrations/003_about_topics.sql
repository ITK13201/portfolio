-- +goose Up
-- +goose StatementBegin
CREATE TABLE about_topics
(
    id             INTEGER AUTO_INCREMENT NOT NULL PRIMARY KEY,
    title          VARCHAR(64)  NOT NULL UNIQUE,
    description_jp VARCHAR(256) NOT NULL,
    description_en VARCHAR(256) NOT NULL DEFAULT '',
    image_id       INTEGER      NOT NULL,
    priority       INTEGER      NOT NULL,
    created_at     DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at     DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_about_topics_image_id
        FOREIGN KEY (image_id)
            REFERENCES images (id)
            ON DELETE RESTRICT ON UPDATE RESTRICT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS about_topics;
-- +goose StatementEnd
