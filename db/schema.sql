CREATE TABLE subscriptions (
    id integer NOT NULL,
    name character varying(120) NOT NULL,
    user_id integer NOT NULL,
    price integer NOT NULL,
    url varchar(120) NOT NULL
);

ALTER TABLE ONLY subscriptions
    ADD CONSTRAINT subscriptions_pkey PRIMARY KEY (id);

ALTER TABLE ONLY subscriptions ALTER COLUMN id SET DEFAULT nextval('subscriptions_id_seq'::regclass);


CREATE TABLE users (
    id integer NOT NULL,
    username character varying(75) NOT NULL
);

ALTER TABLE ONLY users ALTER COLUMN id SET DEFAULT nextval('users_id_seq'::regclass);

ALTER TABLE ONLY users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);

ALTER TABLE ONLY users
    ADD CONSTRAINT users_username_key UNIQUE (username);

ALTER TABLE ONLY subscriptions
    ADD CONSTRAINT subscriptions_user_id_fkey FOREIGN KEY (user_id) REFERENCES users(id);
