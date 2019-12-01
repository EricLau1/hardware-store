create table if not exists categories(
id bigint primary key auto_increment,
description varchar(256) not null unique,
created_at timestamp default current_timestamp,
updated_at timestamp default current_timestamp
)
engine = InnoDB
default charset = utf8;

create table if not exists products(
id bigint primary key auto_increment,
name varchar(512) not null unique,
price decimal(10,2) default 0.0,
quantity int(10) unsigned default 0,
status char(1) default 0,
category_id bigint not null,
created_at timestamp default current_timestamp,
updated_at timestamp default current_timestamp,
constraint products_category_id foreign key(category_id) references categories(id)
on delete cascade on update cascade
)
engine = InnoDB
default charset = utf8;