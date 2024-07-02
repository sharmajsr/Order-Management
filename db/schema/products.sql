CREATE TABLE IF NOT EXISTS public.products (
    id bigInt PRIMARY KEY,
    price BIGINT NOT NULL,
    details TEXT NOT NULL,
    quantity_available INT NOT NULL
);