create table if not exists public.orders
(
    id serial not null
        constraint order_pk
            primary key,
    item text not null,
    quantity integer not null
);