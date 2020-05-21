create table products
(
    id          SERIAL PRIMARY KEY,
    name        VARCHAR            DEFAULT '',
    description VARCHAR            DEFAULT '',
    create_at   timestamp NOT NULL DEFAULT (now() at time zone 'utc'),
    updated_at  timestamp NOT NULL DEFAULT (now() at time zone 'utc')
)
