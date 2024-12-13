create table if not exists logs(
    uuid uuid not null default gen_random_uuid(),
    "path" text,
    method varchar(255),
    status_code int,
    interval numeric(3,3),
    body text,
    user_uuid uuid,
    created_at timestamp default current_timestamp,
    
    primary key(uuid)
)