
db:
	docker run --name test --rm -p 5432:5432 -e POSTGRES_PASSWORD=root -e POSTGRES_DB=db -d postgres
