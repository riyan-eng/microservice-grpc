create table
  if not exists notification_type (
    id uuid primary key,
    code varchar(50) not null,
    "name" varchar(50) not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    constraint u_notification_type_code unique (code)
  );