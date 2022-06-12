include .env
export

postgres:
  docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=$(POSTGRES_USER) -e POSTGRES_PASSWORD=$(POSTGRES_PASSWORD) -d postgres

createdb:
  docker exec -it postgres12 createdb --username=$(POSTGRES_USER) --owner=$(POSTGRES_USER) $(POSTGRES_DB)

dropdb:
  docker exec -it postgres12 dropdb $(POSTGRES_DB)

migrate:
	migrate -database postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@localhost:5432/$(POSTGRES_DB)?sslmode=disable -path db/migrations up

.PHONY: postgres createdb dropdb migrate