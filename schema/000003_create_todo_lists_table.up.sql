create table todo_lists
(
    id          int          not null unique,
    title       varchar(255) not null,
    description varchar(255) not null,
    primary key (id)
);