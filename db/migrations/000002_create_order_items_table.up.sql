CREATE TABLE IF NOT EXISTS order_items(
   id serial PRIMARY KEY,
   order_id INTEGER NOT NULL,
   product_id INTEGER NOT NUll,
   product_name VARCHAR (50) NOT NULL,
   product_price NUMERIC(10, 2) NOT NULL,
   quantity INTEGER NOT NULL
);