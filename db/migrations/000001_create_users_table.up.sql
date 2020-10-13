CREATE TABLE users (
    id serial PRIMARY KEY,
    username varchar(75) UNIQUE NOT NULL
)
