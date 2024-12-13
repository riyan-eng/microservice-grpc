create table "permissions"(
    uuid uuid not null default gen_random_uuid(),
    parent uuid,
    code varchar(255),
    "name" varchar(255),
    full_method varchar(255),
    deleted_at timestamp,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    
    primary key(uuid),
    unique(code),
    unique(full_method)
);