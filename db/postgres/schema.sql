create table users(
    id integer primary key not null
);
	
create table segments(
	id serial primary key,
	name varchar not null unique
);
	
create table history(
	id serial primary key,
	user_id integer not null,
	segment_name varchar not null,
	operation varchar not null,
	operation_time timestamp default now(),
	foreign key (user_id) references users(id),
);
	
create table follows(
	user_id integer not null,
	segment_id integer not null,
	expire timestamp default null,
	unique (user_id, segment_id),
	foreign key (user_id) references users (id) on delete cascade,
	foreign key (segment_id) references segments (id) on delete cascade
);