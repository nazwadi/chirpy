-- +goose UP
CREATE TABLE users (
  id UUID NOT NULL PRIMARY KEY,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  email text NOT NULL
);

-- +goose Down
DROP TABLE users;
