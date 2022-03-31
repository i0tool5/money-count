DOCKER_COMPOSE=docker-compose -f docker/docker-compose.yaml
PAYMENTS_DB=postgres://tempuser:temppass@localhost:5432/tempdb?sslmode=disable

.PHONY: docker.run docker.stop db.create db.drop migrate.up migrate.down

docker.run:
	${DOCKER_COMPOSE} up -d

docker.stop:
	${DOCKER_COMPOSE} down

db.create:
	docker exec money-count-postgres psql -h 127.0.0.1 -U postgres \
		-f etc/db/scripts/create_db.sql

db.drop:
	docker exec money-count-postgres psql -h 127.0.0.1 -U postgres \
		-f etc/db/scripts/drop_db.sql

# migrate.up:
# 	docker run --mount=type=bind,src=${pwd}/docker/etc/migrations,dst=migrations/ --network host migrate/migrate \
# 	-path =/migrations/ -database ${PAYMENTS_DB} up

migrate.up:
	docker run -v ${pwd}/docker/etc/db/migrations:/migrations --network host migrate/migrate \
    -path=/migrations/ -database ${PAYMENTS_DB} up 2

migrate.down:
	docker run --mount=type=bind,src=${pwd}/docker/etc/migrations:/migrations/ --network host migrate/migrate \
	-path =/migrations/ -database ${PAYMENTS_DB} down




