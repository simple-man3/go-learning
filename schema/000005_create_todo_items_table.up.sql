create table todo_items
(
    id          int          not null unique,
    title       varchar(255) not null,
    description varchar(255),
    done        bit          not null default false
);