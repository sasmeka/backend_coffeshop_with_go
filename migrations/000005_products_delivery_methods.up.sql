CREATE TABLE public.products_delivery_methods (
	id_product_delivery_method uuid NOT NULL DEFAULT gen_random_uuid(),
	id_product uuid NOT NULL,
	id_dm uuid NOT NULL,
	create_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	update_at timestamp NULL,
	CONSTRAINT products_delivery_methods_pkey PRIMARY KEY (id_product_delivery_method),
	CONSTRAINT products_delivery_methods_id_dm_fkey FOREIGN KEY (id_dm) REFERENCES public.delivery_methods(id_dm),
	CONSTRAINT products_delivery_methods_id_product_fkey FOREIGN KEY (id_product) REFERENCES public.products(id_product)
);