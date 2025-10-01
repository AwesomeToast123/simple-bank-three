postgres:
	docker run --name postgres_redux -p 5432:5432 -e POSTGRES_USER=user -e POSTGRES_PASSWORD=mysecret -d postgres:12-alpine

createdb:
	docker exec -it postgres_redux createdb -U user simple_bank_three

dropdb:
	docker exec -it postgres_redux dropdb -U user simple_bank_three

migrateup:
	migrate -path db/migration -database "postgresql://user:mysecret@localhost:5432/simple_bank_three?sslmode=disable" force 1 -verbose up 

migratedown:
	migrate -path db/migration -database "postgresql://user:mysecret@localhost:5432/simple_bank_three?sslmode=disable" -verbose down 
	
sqlc:
	sqlc generate
	
test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test
