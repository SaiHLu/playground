run_unix:
	air -c .air.unix.conf

run_windows:
	air -c .air.windows.conf

migration_create:
	migrate create -ext sql -dir ./database/migrations -seq $(NAME)

migration_up: 
	migrate -path database/migrations/ -database "postgres://root:root@localhost:5432/fiber-demo?sslmode=disable" -verbose up

migration_down: 
	migrate -path database/migrations/ -database "postgres://root:root@localhost:5432/fiber-demo?sslmode=disable" -verbose down

migration_fix: 
	migrate -path database/migrations/ -database "postgres://root:root@localhost:5432/fiber-demo?sslmode=disable" force $(VERSION)