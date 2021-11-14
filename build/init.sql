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

insert into users values ('f3bf75a9-ea4c-4f57-9161-cfa8f96e2d0b', 'admin@mail.ru', '$2a$04$Q65Ug0F8llqw4DOPZ13gu.iWR7pDu7zwEdg9SxElmdUQoKnpD2CGe', 'Ivan', 'Ivanov', true, 'test', 'Moscow');

create table accounts
(
    id            varchar(36)  not null,
    user_id       varchar(36)  not null,
    account_id    varchar(36)  not null,
    currency_code varchar(256) null,
    primary key (user_id, account_id),
    constraint user_accounts_id_uindex
        unique (id),
    constraint user_accounts_users_id_fk
        foreign key (user_id) references users (id)
);

create index user_accounts_accounts_id_fk
    on accounts (account_id);

INSERT INTO trade.accounts (id, user_id, account_id, currency_code) VALUES ('a02bc569-8de4-496b-ad19-984bbcf2ed26', 'f3bf75a9-ea4c-4f57-9161-cfa8f96e2d0b', 'd84db1a4-4f6c-4871-9ffa-1b3564c44111', 'RU');