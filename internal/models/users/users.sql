CREATE TABLE users (
    id        bigserial PRIMARY KEY,
    username  text NOT NULL UNIQUE,
    firstname text,
    lastname  text,
    password  text
);
