CREATE TABLE public.sizes (
	id_size uuid NOT NULL DEFAULT gen_random_uuid(),
	name_size varchar(255) NOT NULL,
	abbreviation varchar(3) NOT NULL,
	create_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	update_at timestamp NULL,
	CONSTRAINT sizes_pkey PRIMARY KEY (id_size)
);