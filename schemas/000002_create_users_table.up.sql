create table users
(
    id            int          auto_increment not null,
    name          varchar(255) not null,
    username      varchar(255) not null unique,
    password_hash varchar(255) not null,
    primary key (id)
);