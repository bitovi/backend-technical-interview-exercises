do $$
begin
	for i in 1..50 loop
		insert into "order" (customer, products) values
		(
			floor(random() * 50) + 1,
			array(select trunc(random() * 6 + 1)::integer from generate_series(1, trunc(random() * 5 + 1)::integer))
		);
	end loop;
end; $$;