migration_up: 
	migrate -path database/migrations/ -database "postgres://root:root@localhost:5432/fiber-demo?sslmode=disable" -verbose up

migration_down: 
	migrate -path database/migrations/ -database "postgres://root:root@localhost:5432/fiber-demo?sslmode=disable" -verbose down

migration_fix: 
	migrate -path database/migrations/ -database "postgres://root:root@localhost:5432/fiber-demo?sslmode=disable" force VERSION