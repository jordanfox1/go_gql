regen:
	go run github.com/99designs/gqlgen generate --verbose

psql:
	sudo -u postgres psql

reset-db:
	migrate -path "postgres/migrations" -database postgres://postgres:password@localhost:5432/meetmeup_dev?sslmode=disable drop
	migrate -path "postgres/migrations" -database postgres://postgres:password@localhost:5432/meetmeup_dev?sslmode=disable up