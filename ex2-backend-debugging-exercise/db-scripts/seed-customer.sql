CREATE OR REPLACE FUNCTION random_string(length INTEGER) RETURNS TEXT AS $$
DECLARE
    characters TEXT := 'ABCDEFGHIJKLMNOPQRSTUVWXYZ';
    result TEXT := '';
		len int := 0;
BEGIN
    FOR i in 1..length LOOP
        result := result || substring(characters, floor(random() * length(characters))::int + 1, 1);
    END LOOP;
    RETURN result;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION get_customer_details()
RETURNS TABLE(customer_id INT, first_name VARCHAR, last_name VARCHAR, email VARCHAR)
LANGUAGE plpgsql
AS $$
BEGIN
    RETURN QUERY
    SELECT c.customer_id, cn.first_name::VARCHAR(180), cn.last_name::VARCHAR(180), ce.email::VARCHAR(18)
    FROM public.customer c
    JOIN public.customer_email ce 
    ON c.email_id = ce.email_id
    JOIN public.customer_name cn 
    ON c.name_id = cn.name_id;
END;
$$;

DO $$
DECLARE
    i INTEGER := 0;
BEGIN
    FOR i IN 1..50 LOOP
        INSERT INTO customer_name (first_name, last_name)
        VALUES (random_string(6), random_string(8));
    END LOOP;
END;
$$;

DO $$
DECLARE
    domain TEXT;
    domains TEXT[] := ARRAY['gmail.com', 'yahoo.com', 'outlook.com', 'hotmail.com', 'icloud.com'];
BEGIN
    FOR i IN 1..50 LOOP
        domain := domains[floor(random()*array_length(domains, 1) + 1)];
        INSERT INTO customer_email (email)
        VALUES (lower(random_string(8) || '@' || domain));
    END LOOP;
END;
$$;

do $$
begin
	for i in 1..50 loop
		insert into "customer" (name_id, email_id)
		VALUES (trunc(random() * 50 + 1)::integer, trunc(random() * 50 + 1)::integer);
	end loop;
end; $$;