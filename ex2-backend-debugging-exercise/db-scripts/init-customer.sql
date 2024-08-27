CREATE TABLE IF NOT EXISTS customer (
  customer_id SERIAL PRIMARY KEY,
	name_id integer,
	email_id integer
);

CREATE TABLE IF NOT EXISTS customer_name (
  name_id SERIAL PRIMARY KEY,
	first_name VARCHAR(180),
	last_name VARCHAR(180)
);

CREATE TABLE IF NOT EXISTS customer_email (
  email_id SERIAL PRIMARY KEY,
	email VARCHAR(180)
);