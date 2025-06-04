-- +goose Up
-- +goose StatementBegin
create extension if not exists "uuid-ossp";

create table if not exists products
(
    uuid UUID primary key,
    store_cost decimal not null,
    store_amount decimal not null check (store_amount > 0)
    );

create type supply_status as enum ('created', 'in_work','served', 'on_the_road', 'shipped', 'done');

create table if not exists supply
(
    uuid UUID primary key default uuid_generate_v4(),
    comment varchar(255),
    creation_date timestamp not null default now(),
    desired_date timestamp not null,
    status supply_status default 'created' not null,
    responsible_user varchar(24) not null,
    edited boolean default false,
    edited_date timestamp,
    cost decimal not null
    );

create table supply_product
(
    product_uuid uuid references products(uuid),
    supply_uuid uuid references supply(uuid),
    amount decimal not null
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

drop type if exists supply_status;

drop table if exists products;

drop table if exists supply;

drop table if exists supply_product;

-- +goose StatementEnd
