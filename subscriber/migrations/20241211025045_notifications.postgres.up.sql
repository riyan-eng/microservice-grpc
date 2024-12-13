create table
  if not exists notifications (
    id uuid primary key,
    "status" varchar(50) not null,
    "type" varchar(50) not null,
    body varchar,
    error varchar,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    constraint fk_notifications_status foreign key ("status") references notification_status (code) on delete no action,
    constraint fk_notifications_type foreign key ("type") references notification_type (code) on delete no action
  );