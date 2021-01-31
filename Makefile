build:
	go build -o bin/main main.go

run:
	MODE=production ./bin/main

db-seed:
	go run cmd/seed_db/seed.go

db-drop:
	go run cmd/drop_db/drop.go
