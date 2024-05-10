create table user_lists
(
    id      int not null unique,
    user_id int,
    list_id int,
    primary key (id),
    foreign key (user_id) references users (id) on delete cascade,
    foreign key (list_id) references todo_lists (id) on delete cascade
);