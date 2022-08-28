create table user(
    user_id uuid DEFAULT uuid_generate_v1 (),
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

create table group(
    group_id uuid DEFAULT uuid_generate_v1 (),
    group_name varchar not null,
    teachers uuid[],
    students uuid[],
    created_on date default current_timestamp,
    updated_on date default current_timestamp,
    primary key(group_id)
);

create table test(
    test_id uuid DEFAULT uuid_generate_v1(),
    title varchar not null,
    description text,
    created_by uuid foreign key references user(user_id),
    enrolled uuid[],
    schedule date,
    start_time time,
    end_time time,
    link varchar not null,
    created_on timestamp default current_timestamp,
    updated_on timestamp default current_timestamp,
    primary key(test_id)
);

create table score(
    test_id uuid foreign key references test(test_id),
    created_by uuid foreign key references user(user_id),
    test_taker uuid foreign key references user(user_id),
    checked_by uuid foreign key references user(user_id),
    checked_on timestamp default current_timestamp,
    status varchar default 'pass',
    primary key(test_id, test_taker)
)