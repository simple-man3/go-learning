# create table todo_items
# (
#     id          int          not null unique,
#     title       varchar(255) not null,
#     description varchar(255),
#     done        bit          not null default false,
# );
#
# create table lists_items
# (
#     id      int not null unique,
#     item_id int not null,
#     list_id int not null,
#
#     foreign key (item_id) references todo_items on delete cascade,
#     foreign key (list_id) references todo_lists on delete cascade,
# );