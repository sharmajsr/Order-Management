create table if not exists users(
    id serial primary key,
    name varchar(256) not null,
    username varchar(256) not null,
    password varchar(256) not null
);