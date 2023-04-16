

run:
	go run cmd/lamoda/main.go
test:
	go test ./... -cover -count=1
generate:
	go generate ./...

d-up:
	docker run --name=postgres -e POSTGRES_PASSWORD='qwerty' -p 5434:5432 -d --rm postgres

d-stop:
	docker stop postgres

d-exec:
	docker exec -it postgres /bin/bash

m-up:
	migrate -path ./scheme -database 'postgres://postgres:qwerty@localhost:5434/postgres?sslmode=disable' up

m-down:
	migrate -path ./scheme -database 'postgres://postgres:qwerty@localhost:5434/postgres?sslmode=disable' down
