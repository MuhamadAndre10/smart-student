POSTGRES_URL = postgres://postgres:root@localhost:5432/smartstudentsdb?sslmode=disable

migrate-up:
	@echo "Migration up..."
	migrate -path db/migrations -database ${POSTGRES_URL} up
	@echo "Migration up done"

migrate-down:
	@echo "Migration down..."
	migrate -path db/migrations -database ${POSTGRES_URL} down
	@echo "Migration down done"
