CREATE TABLE IF NOT EXISTS users(
    id UUID DEFAULT uuid_generate_v4(),
    username    VARCHAR NOT NULL,
    email   VARCHAR NOT NULL,
    password VARCHAR NOT NULL,
    created_at  TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,

    PRIMARY KEY(id)
)