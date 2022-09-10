create extension if not exists "uuid-ossp";

create table users (
    user_id uuid DEFAULT uuid_generate_v1(),
    first_name varchar not null,
    last_name varchar,
    email varchar unique,
    password varchar not null,
    profile_pic varchar,
    created_on date default current_timestamp,
    updated_on date default current_timestamp,
    role varchar not null,
    primary key(user_id)
);

create table groups (
    group_id uuid DEFAULT uuid_generate_v1 (),
    group_name varchar not null,
    teachers uuid[],
    students uuid[],
    created_on date default current_timestamp,
    updated_on date default current_timestamp,
    primary key(group_id)
);

create table tests (
    test_id uuid DEFAULT uuid_generate_v1(),
    title varchar not null,
    description text,
    created_by uuid,
    enrolled uuid[],
    schedule date,
    start_time time,
    end_time time,
    link varchar not null,
    created_on timestamp default current_timestamp,
    updated_on timestamp default current_timestamp,
    primary key(test_id),
    constraint fk_tests_created_by foreign key (created_by) references users(user_id)
);

create table scores (
    test_id uuid,
    test_taker uuid,
    checked_by uuid,
    checked_on timestamp default current_timestamp,
    status varchar default 'pass',
    primary key(test_id, test_taker),
    constraint fk_scores_test_id foreign key (test_id) references tests(test_id),
    constraint fk_scores_checked_by foreign key (checked_by) references users(user_id),
    constraint fk_scores_test_taker foreign key (test_taker) references users(user_id)
);