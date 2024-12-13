create table user_datas(
    uuid uuid not null default gen_random_uuid(),
    user_uuid uuid,
    role_code varchar(255),
    is_active boolean not null default true,
    deleted_at timestamp,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    
    primary key(uuid),
    unique(user_uuid, role_code),
    constraint fk_user foreign key(user_uuid) references users(uuid) on delete cascade,
    constraint fk_role foreign key(role_code) references roles(code) on delete no action
);