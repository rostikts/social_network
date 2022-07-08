create table if not exists users (
    id serial primary key,
    username varchar unique not null,
    first_name varchar not null,
    last_name varchar not null,
    email varchar not null,
    password varchar not null
)