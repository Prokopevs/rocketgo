create table if not exists users(
  id serial not null,
  firstname varchar not null,
  username varchar,
  created_at date not null
);