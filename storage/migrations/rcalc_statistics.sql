CREATE TABLE public.rcalc_statistics (
	id BIGSERIAL primary KEY,
    age int8 NULL,
    lifespan int8 NULL,
    income float8 NULL,
    expenses float8 NULL,
    inflation float8 NULL,
    currency varchar NULL,
    raise float8 NULL,
    advanced_mode bool NULL,
    investments float8 NULL,
    returns float8 NULL,
    other_expenses int8 NULL,
    current_savings float8 NULL,
    created_at timestamp NULL
);