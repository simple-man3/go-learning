create table users
(
    id            int          not null,
    name          varchar(255) not null,
    username      varchar(255) not null unique,
    password_hash varchar(255) not null,
    primary key (id)
)

create table todo_lists
(
    id          int          not null unique,
    title       varchar(255) not null,
    description varchar(255) not null,
    primary key (id)
)

create table users_lists
(
    id      int not null unique,
    user_id int,
    list_id int,
    primary key (id),
    foreign key (user_id) references users (id) on delete cascade,
    foreign key (list_id) references todo_lists (id) on delete cascade
)

create table todo_items
(
    id          int          not null unique,
    title       varchar(255) not null,
    description varchar(255),
    done        bit          not null default false,
)

create table lists_items
(
    id      int not null unique,
    item_id int not null,
    list_id int not null,

    foreign key (item_id) references todo_items on delete cascade,
    foreign key (list_id) references todo_lists on delete cascade,
)