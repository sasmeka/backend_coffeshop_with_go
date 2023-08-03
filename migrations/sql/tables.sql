CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE public.users (
	id_user uuid primary key NULL DEFAULT uuid_generate_v4(),
  	displayname varchar(255) NOT NULL,
	first_name text NOT NULL,
	last_name text NOT NULL,
    gender varchar(255) NOT NULL DEFAULT 'male'::character varying,
	phone varchar(15) NULL,
	email text NOT NULL,
	pass text NOT NULL,
  	birth_date timestamp,
	status_verification bpchar(1) NOT NULL DEFAULT '0'::bpchar,
	"role" varchar(6) NOT NULL DEFAULT 'user'::character varying,
	image text NOT NULL DEFAULT '/static/img/Default_Profile.png'::text,
	create_at timestamp not null DEFAULT CURRENT_TIMESTAMP,
	update_at timestamp
);

CREATE TABLE public.products (
	id_product uuid primary key NULL DEFAULT uuid_generate_v4(),
	name_product varchar(255) NOT NULL,
	description text,
	stok int NOT NULL DEFAULT 0,
	favorite varchar(1) NOT NULL DEFAULT '0',
	image text NOT NULL DEFAULT '/static/product/Default_Product.png'::text,
	create_at timestamp not null DEFAULT CURRENT_TIMESTAMP,
	update_at timestamp
);

CREATE TABLE public.sizes (
	id_size uuid primary key NULL DEFAULT uuid_generate_v4(),
    name_size varchar(255) NOT NULL,
    abbreviation varchar(3) not null,
	create_at timestamp not null DEFAULT CURRENT_TIMESTAMP,
	update_at timestamp
);

CREATE TABLE public.delivery_methods (
	id_dm uuid primary key NULL DEFAULT uuid_generate_v4(),
    name_dm varchar(255) NOT NULL,
	create_at timestamp not null DEFAULT CURRENT_TIMESTAMP,
	update_at timestamp
);

CREATE TABLE public.products_sizes (
	id_product_size uuid primary key NULL DEFAULT uuid_generate_v4(),
    id_product uuid NOT NULL,
    id_size uuid not null,
    price int not null,
	create_at timestamp not null DEFAULT CURRENT_TIMESTAMP,
	update_at timestamp,
	FOREIGN KEY (id_product) REFERENCES products (id_product),
	FOREIGN KEY (id_size) REFERENCES sizes (id_size)
	);

CREATE TABLE public.products_delivery_methods (
	id_product_delivery_method uuid primary key NULL DEFAULT uuid_generate_v4(),
    id_product uuid NOT NULL,
    id_dm uuid not null,
	create_at timestamp not null DEFAULT CURRENT_TIMESTAMP,
	update_at timestamp,
	FOREIGN KEY (id_product) REFERENCES products (id_product),
	FOREIGN KEY (id_dm) REFERENCES delivery_methods (id_dm)
	);

-- CREATE TABLE public.favorites_products (
-- 	id_favorite uuid primary key NULL DEFAULT uuid_generate_v4(),
--     id_product uuid NOT NULL,
--     id_user uuid NOT NULL,
-- 	create_at timestamp not null DEFAULT CURRENT_TIMESTAMP,
-- 	update_at timestamp,
-- 	FOREIGN KEY (id_product) REFERENCES products (id_product),
-- 	FOREIGN KEY (id_user) REFERENCES users (id_user)
-- );


