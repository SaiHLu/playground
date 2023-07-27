CREATE TABLE IF NOT EXISTS users (
    id  uuid DEFAULT uuid_generate_v4(),
    username    varchar(50) NOT NULL,
    email   varchar(300) UNIQUE NOT NULL,


    PRIMARY KEY(id)
)

