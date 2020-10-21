CREATE TABLE subscriptions (
    id serial PRIMARY KEY,
    user_id integer REFERENCES users (id) NOT NULL,
    name varchar(120) NOT NULL,
    price integer NOT NULL
)
