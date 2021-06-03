create schema if not exists public;

create table if not exists users
(
	id serial constraint users_pkprimary key,
	name varchar,
	age int,
	"from" varchar,
	vip bool,
	created_at timestamp default current_timestamp
);