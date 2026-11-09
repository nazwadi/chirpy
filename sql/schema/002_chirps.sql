-- +goose UP
CREATE TABLE chirps (
    id UUID NOT NULL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    body text NOT NULL,
    user_id UUID NOT NULL
);

-- +goose Down
DROP TABLE chirps;
