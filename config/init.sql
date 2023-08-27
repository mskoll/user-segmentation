create table users
(
    id       serial primary key,
    username varchar(255) not null
);

create table segment
(
    id         serial primary key,
    name       varchar(255) not null,
    percent    int          not null default 0,
    created_at timestamp    not null default now(),
    deleted_at timestamp             default null
);

create table user_segment
(
    id         serial primary key,
    user_id    int       not null,
    segment_id int       not null,
    created_at timestamp not null default now(),
    deleted_at timestamp          default null,
    foreign key (user_id) references users (id),
    foreign key (segment_id) references segment (id)
);

create unique index idx_user_segment on user_segment (user_id, segment_id, created_at);