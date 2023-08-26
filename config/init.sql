create table users
(
    id serial primary key ,
    username varchar
);

create table segment
(
    id serial primary key ,
    name varchar
);

create table user_segment
(
    id serial primary key ,
    user_id int not null ,
    segment_id int not null ,
    foreign key (user_id) references users (id),
    foreign key (segment_id) references segment (id)
)