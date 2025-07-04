local-migration:
	goose -dir migrations postgres "host=localhost port=5432 user=postgres password=postgres dbname=postgres sslmode=disable" up

tear-migration:
	goose -dir migrations postgres "host=localhost port=5432 user=postgres password=postgres dbname=postgres sslmode=disable" down