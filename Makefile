build:
	env GOOS=linux GOARCH=amd64 go build -o bin/main -v .

run:
	MODE=production ./bin/main

db-seed:
	go run cmd/seed_db/seed.go

db-drop:
	go run cmd/drop_db/drop.go
