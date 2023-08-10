CREATE TABLE public.products_sizes (
	id_product_size uuid NOT NULL DEFAULT gen_random_uuid(),
	id_product uuid NOT NULL,
	id_size uuid NOT NULL,
	price int4 NOT NULL,
	create_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	update_at timestamp NULL,
	stok int4 NOT NULL DEFAULT 0,
	CONSTRAINT products_sizes_pkey PRIMARY KEY (id_product_size),
	CONSTRAINT products_sizes_id_product_fkey FOREIGN KEY (id_product) REFERENCES public.products(id_product),
	CONSTRAINT products_sizes_id_size_fkey FOREIGN KEY (id_size) REFERENCES public.sizes(id_size)
);