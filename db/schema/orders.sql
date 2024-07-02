create table if not exists public.orders(
    id bigInt primary key ,
    product_id bigInt not null,
    user_id bigInt not null,
    price bigInt not null,
    status text not null,
    quantity_ordered bigint not null,
    foreign key (product_id) references products(id),
    foreign key (user_id) references users(id)
);