create table users(
    id int primary key not null
);

create table segments(
    id int serial,
    name varchar(32) not null unique
);

create table history(
    id int serial,
    user_id int not null references users (id),
    segment_name varchar(32) not null references segments (name),
    operation varchar not null,
    operationTime timestamp default now()
);

create table follows(
    user_id int not null references users (id) on delete cascade,
    segment_name int not null references segments (name) on delete cascade,
    expire timestamp default now()
);