create table users(
    uuid uuid not null default gen_random_uuid(),
    email varchar(255) not null,
    username varchar(255) not null,
    "password" text,
    photo_url text,
    is_active boolean not null default true,
    deleted_at timestamp,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    
    primary key(uuid),
    unique(email)
);