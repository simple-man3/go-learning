create table list_items
(
    id int not null unique,
    todo_item_id int not null,
    todo_list_id int not null,
    foreign key (todo_item_id) references todo_items (id),
    foreign key (todo_list_id) references todo_lists (id)
)