createdb:
	docker exec -it postgres12 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres12 dropdb simple_bank

migrateup:
	 migrate -path db/migration -database "postgresql://root:root@localhost:5432/simple_bank?sslmode=disable" --verbose up

migratedown:
	 migrate -path db/migration -database "postgresql://root:root@localhost:5432/simple_bank?sslmode=disable" --verbose down	

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go	

mock: 
	mockgen -package mockdb  --destination db/mock/store.go  github.com/wasinaphatlilawatthananan/go-postgres/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server mock