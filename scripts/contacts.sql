CREATE TABLE public.contacts (
	id BIGSERIAL primary KEY,
	"name" varchar NULL,
	email varchar NULL,
	subject varchar NULL,
	message varchar NULL,
	created_at timestamp NULL
);