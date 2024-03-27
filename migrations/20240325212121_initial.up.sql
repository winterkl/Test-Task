-- Стеллажи -- Стеллажи -- Стеллажи -- Стеллажи -- Стеллажи -- Стеллажи
create table if not exists racks
(
    id    serial
        constraint racks_pk
            primary key,
    title varchar(32) not null
);

INSERT INTO racks (id, title) VALUES (1, 'А');
INSERT INTO racks (id, title) VALUES (2, 'Б');
INSERT INTO racks (id, title) VALUES (3, 'Ж');
INSERT INTO racks (id, title) VALUES (4, 'З');
INSERT INTO racks (id, title) VALUES (5, 'В');


-- Категории -- Категории -- Категории -- Категории -- Категории -- Категории --
create table if not exists product_categories
(
    id    serial
        constraint product_categories_pk
            primary key,
    title varchar(100) not null
        constraint product_categories_pk_2
            unique
);

INSERT INTO product_categories (id, title) VALUES (1, 'Ноутбуки');
INSERT INTO product_categories (id, title) VALUES (2, 'Телевизоры');
INSERT INTO product_categories (id, title) VALUES (3, 'Телефоны');
INSERT INTO product_categories (id, title) VALUES (4, 'Компьютеры');
INSERT INTO product_categories (id, title) VALUES (5, 'Часы');
INSERT INTO product_categories (id, title) VALUES (6, 'Переферия');


-- Товары -- Товары -- Товары -- Товары -- Товары -- Товары -- Товары -- Товары --
create table if not exists products
(
    id       serial
        constraint products_pk
            primary key,
    title   varchar(255) not null
        constraint products_pk_2
            unique,
    category_id integer not null
        constraint product_product_categories_id_fk
            references product_categories
);

INSERT INTO products (id, title, category_id) VALUES (1, 'Ноутбук', 1);
INSERT INTO products (id, title, category_id) VALUES (2, 'Телевизор', 2);
INSERT INTO products (id, title, category_id) VALUES (3, 'Телефон', 3);
INSERT INTO products (id, title, category_id) VALUES (4, 'Системный блок', 4);
INSERT INTO products (id, title, category_id) VALUES (5, 'Часы', 5);
INSERT INTO products (id, title, category_id) VALUES (6, 'Микрофон', 6);


-- Заказы -- Заказы -- Заказы -- Заказы -- Заказы -- Заказы -- Заказы --
create table if not exists orders
(
    id         serial
        constraint orders_pk
            primary key,
    created_at timestamp default CURRENT_TIMESTAMP not null
);

INSERT INTO orders (id) VALUES (10);
INSERT INTO orders (id) VALUES (11);
INSERT INTO orders (id) VALUES (14);
INSERT INTO orders (id) VALUES (15);


-- Категории Стеллажей -- Категории Стеллажей -- Категории Стеллажей -- Категории Стеллажей --
create table category_racks
(
    id          serial
        constraint category_rack_pk
            primary key,
    category_id integer not null
        constraint category_rack_product_categories_id_fk
            references product_categories,
    rack_id     integer not null
        constraint category_rack_racks_id_fk
            references racks,
    is_main     boolean not null,
    constraint category_rack_pk_2
        unique (category_id, rack_id)
);

INSERT INTO category_racks (id, category_id, rack_id, is_main) VALUES (1, 1, 1, true);
INSERT INTO category_racks (id, category_id, rack_id, is_main) VALUES (2, 2, 1, true);
INSERT INTO category_racks (id, category_id, rack_id, is_main) VALUES (3, 3, 2, true);
INSERT INTO category_racks (id, category_id, rack_id, is_main) VALUES (4, 3, 4, false);
INSERT INTO category_racks (id, category_id, rack_id, is_main) VALUES (5, 3, 5, false);
INSERT INTO category_racks (id, category_id, rack_id, is_main) VALUES (6, 4, 3, true);
INSERT INTO category_racks (id, category_id, rack_id, is_main) VALUES (7, 5, 3, true);
INSERT INTO category_racks (id, category_id, rack_id, is_main) VALUES (8, 5, 1, false);
INSERT INTO category_racks (id, category_id, rack_id, is_main) VALUES (9, 6, 3, true);


-- Заказы Товаров -- Заказы Товаров -- Заказы Товаров -- Заказы Товаров -- Заказы Товаров --
create table if not exists order_products
(
    id         serial
        constraint order_product_pk
            primary key,
    order_id   integer  not null
        constraint order_product_orders_id_fk
            references orders,
    product_id integer  not null
        constraint order_product_product_id_fk
            references products,
    count      smallint not null,
    rack_id    integer  not null
        constraint order_product_racks_id_fk
            references racks,
    constraint order_product_pk_2
        unique (order_id, product_id)
);

INSERT INTO order_products (id, order_id, product_id, count, rack_id) VALUES (1, 10, 1, 2, 1);
INSERT INTO order_products (id, order_id, product_id, count, rack_id) VALUES (2, 10, 3, 1, 2);
INSERT INTO order_products (id, order_id, product_id, count, rack_id) VALUES (3, 10, 6, 1, 3);
INSERT INTO order_products (id, order_id, product_id, count, rack_id) VALUES (4, 11, 2, 3, 1);
INSERT INTO order_products (id, order_id, product_id, count, rack_id) VALUES (5, 14, 1, 3, 1);
INSERT INTO order_products (id, order_id, product_id, count, rack_id) VALUES (6, 14, 4, 4, 3);
INSERT INTO order_products (id, order_id, product_id, count, rack_id) VALUES (7, 15, 5, 1, 3);