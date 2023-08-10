CREATE TABLE public.products (
	id_product uuid NOT NULL DEFAULT gen_random_uuid(),
	name_product varchar(255) NOT NULL,
	description text NULL,
	favorite varchar(1) NOT NULL DEFAULT '0'::character varying,
	image text NOT NULL DEFAULT '/static/product/Default_Product.png'::text,
	create_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	update_at timestamp NULL,
	CONSTRAINT products_pkey PRIMARY KEY (id_product)
);