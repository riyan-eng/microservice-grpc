create table example(
    id serial not null,
    uuid uuid not null default gen_random_uuid(),
    "name" varchar(255),
    detail text,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    
    primary key(id),
    unique(uuid),
    unique("name")
);