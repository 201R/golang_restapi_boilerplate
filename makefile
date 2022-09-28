
db:
	docker run --name db --rm -p 5432:5432 -e POSTGRES_PASSWORD=password -e POSTGRES_DB=db -d postgres
