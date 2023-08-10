CREATE TABLE public.users (
	id_user uuid NOT NULL DEFAULT gen_random_uuid(),
	displayname varchar(255) NULL,
	first_name text NULL,
	last_name text NULL,
	gender varchar(255) NOT NULL DEFAULT 'male'::character varying,
	phone varchar(15) NULL,
	email text NOT NULL,
	pass text NOT NULL,
	birth_date timestamp NULL,
	status_verification bpchar(1) NOT NULL DEFAULT '0'::bpchar,
	"role" varchar(6) NOT NULL DEFAULT 'user'::character varying,
	image text NOT NULL DEFAULT '/static/img/Default_Profile.png'::text,
	create_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	update_at timestamp NULL,
	CONSTRAINT users_pkey PRIMARY KEY (id_user)
);