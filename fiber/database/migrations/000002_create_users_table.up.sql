CREATE TABLE IF NOT EXISTS "users" (
    "id"  UUID NOT NULL DEFAULT (uuid_generate_v4()),
    "username"    VARCHAR NOT NULL,
    "email"   VARCHAR NOT NULL UNIQUE,
    "password"  VARCHAR NOT NULL,

    CONSTRAINT  "users_pkey"    PRIMARY KEY("id")
)

