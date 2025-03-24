create table product (
	id serial primary key,
	product_name varchar(50) not null,
	price numeric(10, 2) not null
);

insert into product (product_name, price) values ('Prod 001', 100);
insert into product (product_name, price) values ('Prod 002', 200);
insert into product (product_name, price) values ('Prod 003', 300);
insert into product (product_name, price) values ('Prod 004', 400);
insert into product (product_name, price) values ('Prod 005', 500);