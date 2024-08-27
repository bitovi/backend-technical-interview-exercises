CREATE TABLE IF NOT EXISTS "order" (
  order_id SERIAL PRIMARY KEY,
  customer integer,
  products integer[]
);