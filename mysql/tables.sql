use imdsb;

create table if not exists files
(
    id         bigint unsigned auto_increment
        primary key,
    created_at datetime(3)     null,
    updated_at datetime(3)     null,
    deleted_at datetime(3)     null,
    file_name  longtext        not null,
    file_type  longtext        not null,
    file_path  longtext        not null,
    room_id    bigint unsigned not null,
    sender_id  bigint unsigned not null
);

create index idx_files_deleted_at
    on files (deleted_at);

create table if not exists messages
(
    id         bigint unsigned auto_increment
        primary key,
    created_at datetime(3)     null,
    updated_at datetime(3)     null,
    deleted_at datetime(3)     null,
    sender_id  bigint unsigned not null,
    room_id    bigint unsigned not null,
    type       longtext        not null,
    content    text            null
);

create index idx_messages_deleted_at
    on messages (deleted_at);

create table if not exists online_users
(
    user_id    bigint unsigned                     not null,
    login_time timestamp default CURRENT_TIMESTAMP not null,
    primary key (user_id)
)
    engine = MEMORY;

create table if not exists rooms
(
    id          bigint unsigned auto_increment
        primary key,
    created_at  datetime(3) null,
    updated_at  datetime(3) null,
    deleted_at  datetime(3) null,
    name        longtext    not null,
    description longtext    null
);

create index idx_rooms_deleted_at
    on rooms (deleted_at);

create table if not exists user_infos
(
    id         bigint unsigned auto_increment
        primary key,
    created_at datetime(3)     null,
    updated_at datetime(3)     null,
    deleted_at datetime(3)     null,
    user_id    bigint unsigned null,
    username   varchar(191)    null,
    nike_name  longtext        null,
    avatar     longtext        null,
    sexuality  longtext        null,
    year       bigint unsigned null,
    month      bigint unsigned null,
    day        bigint unsigned null,
    country    longtext        null,
    province   longtext        null,
    city       longtext        null,
    email      longtext        null,
    constraint uni_user_infos_username
        unique (username)
);

create index idx_user_infos_deleted_at
    on user_infos (deleted_at);

create index idx_user_infos_user_id
    on user_infos (user_id);

create table if not exists user_rooms
(
    id         bigint unsigned auto_increment
        primary key,
    created_at datetime(3)     null,
    updated_at datetime(3)     null,
    deleted_at datetime(3)     null,
    room_id    bigint unsigned not null,
    user_id    bigint unsigned not null,
    role       longtext        null
);

create index idx_user_rooms_deleted_at
    on user_rooms (deleted_at);

create index idx_user_rooms_room_id
    on user_rooms (room_id);

create index idx_user_rooms_user_id
    on user_rooms (user_id);

create table if not exists users
(
    id         bigint unsigned auto_increment
        primary key,
    created_at datetime(3)  null,
    updated_at datetime(3)  null,
    deleted_at datetime(3)  null,
    username   varchar(191) null,
    password   longtext     null,
    constraint uni_users_username
        unique (username)
);

create index idx_users_deleted_at
    on users (deleted_at);


