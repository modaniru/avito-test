create table users(
    id integer primary key not null
);
	
create table segments(
	id integer primary key autoincrement,
	name varchar not null unique
);
	
create table history(
	id integer primary key autoincrement,
	user_id integer not null,
	segment_name varchar not null,
	operation varchar not null,
	operation_time timestamp default CURRENT_TIMESTAMP
);
	
create table follows(
	user_id integer not null,
	segment_id integer not null,
	expire timestamp default null,
	unique (user_id, segment_id),
	foreign key (user_id) references users (id) on delete cascade,
	foreign key (segment_id) references segments (id) on delete cascade
);