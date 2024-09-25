-- +goose Up
CREATE TABLE songs(
    id serial primary key,
    group VARCHAR(256) NOT NULL,
    song VARCHAR(256) NOT NULL,
    link VARCHAR(256),
    text TEXT,
    releaseDate TIMESTAMP NOT NULL DEFAULT now()
);

-- +goose Down
DROP TABLE songs;
