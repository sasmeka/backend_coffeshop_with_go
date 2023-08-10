CREATE TABLE public.delivery_methods (
	id_dm uuid NOT NULL DEFAULT gen_random_uuid(),
	name_dm varchar(255) NOT NULL,
	create_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	update_at timestamp NULL,
	CONSTRAINT delivery_methods_pkey PRIMARY KEY (id_dm)
);