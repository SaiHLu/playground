include .env

run_mac:
	~/go/bin/air

run_windows:
	air -c .air.toml.windows

create_migration:
	migrate create -ext sql -dir database/migrations $(tableName)

run_migrations:
	migrate -database ${POSTGRESQL_URL} -path database/migrations up

revert_migration:
	migrate -database ${POSTGRESQL_URL} -path database/migrations down $(step)

force_migration:
	migrate -database ${POSTGRESQL_URL} -path database/migrations force $(version)