
create table user_credential (
    id varchar(100) primary key,
    email varchar(100) unique not null,
    password varchar(100) not null,
    name varchar(50) not null,
    is_active BOOLEAN
);

create table payment(
    id varchar(100) primary key,
    id_user varchar(100),
    name varchar(100),
    purchase_date date,
    price varchar(100),
    foreign key(id_user) references user_credential(id),
);
