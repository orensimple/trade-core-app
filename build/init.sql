CREATE DATABASE trade;

USE trade;

create table parameters (
    id integer primary key, 
    funds integer, 
    btc integer
);

insert into parameters values (1, 10000, 10);

create table users (
    id VARCHAR(36) primary key,
    email VARCHAR(64) not null,
    password VARCHAR(64) not null,
    first_name VARCHAR(32) not null,
    last_name VARCHAR(32) not null,
    male bool,
    about VARCHAR(256),
    address VARCHAR(256)
);

CREATE INDEX user_first_name_last_name_idx ON users ( first_name, last_name );

CREATE INDEX user_email_idx ON users (email);

create table card_brands (
    brand VARCHAR(36) primary key
);

create table cards (
    card_id VARCHAR(36) primary key,
    brand text references card_brands(brand) on update cascade
);

create table orders (
    order_id VARCHAR(36) primary key,
    user_id VARCHAR(36) references users(id)
);

create table payments (
    order_id VARCHAR(36) primary key references orders(order_id),
    card_id VARCHAR(36) references cards(card_id)
);

insert into users values ('f3bf75a9-ea4c-4f57-9161-cfa8f96e2d0b', 'admin@mail.ru', '$2a$04$Q65Ug0F8llqw4DOPZ13gu.iWR7pDu7zwEdg9SxElmdUQoKnpD2CGe', 'Ivan', 'Ivanov', true, 'test', 'Moscow');

insert into card_brands values ('VISA'), ('AMEX');

insert into cards values ('3224ebc0-0a6e-4e22-9ce8-c6564a1bb6a1', 'VISA');

insert into orders values ('722b694c-984c-4208-bddd-796553cf83e1', 'f3bf75a9-ea4c-4f57-9161-cfa8f96e2d0b');

insert into payments values ('722b694c-984c-4208-bddd-796553cf83e1', '3224ebc0-0a6e-4e22-9ce8-c6564a1bb6a1');