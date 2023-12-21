include .env

DATABASE_URL="postgres://${PG_USER}:${PG_PASS}@0.0.0.0:${PG_PORT}/${PG_BASE}?sslmode=${PG_SSL_MODE}"

run:
	docker-compose -f docker-compose.yaml up

migrate-up:
	migrate -path ./migration -database ${DATABASE_URL} up

migrate-down:
	echo y | migrate -path ./migration -database ${DATABASE_URL} down