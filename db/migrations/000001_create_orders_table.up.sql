CREATE TABLE IF NOT EXISTS orders(
   id serial PRIMARY KEY,
   status VARCHAR (50) NOT NULL,
   total_price NUMERIC(10, 2) NOT NULL,
   paid_value NUMERIC(10, 2) NOT NULL
);