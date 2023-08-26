create table users
(
    id       serial primary key,
    username varchar(255) not null
);

create table segment
(
    id      serial primary key,
    name    varchar(255) not null
);

create table user_segment
(
    id         serial primary key,
    user_id    int not null,
    segment_id int not null,
    foreign key (user_id) references users (id),
    foreign key (segment_id) references segment (id) on delete cascade
);

create table operations
(
    id         serial primary key,
    user_id    int          not null,
    segment_id int          not null,
    type       varchar(255) not null,
    created_at timestamp    not null default now(),
    foreign key (user_id) references users (id)
)