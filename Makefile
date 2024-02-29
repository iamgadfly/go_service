# note: call scripts from /scripts

.PHONY: migrate migrate_down migrate_up migrate_version

# ==============================================================================
# Go migrate mysql

run:
	go run cmd/api/main.go

force:
	migrate -database "postgres://postgres:@tcp(localhost:3306)/go_api" -path migrations  force 1

version:
	#migrate -database postgres://postgres:postgres@localhost:5432/auth_db?sslmode=disable -path migrations version
	migrate -path migrations -database "postgres://postgres:@localhost:5432/go_service?sslmode=disable" version

migrate_up:
	migrate -database "postgres://postgres:@localhost:5432/go_service?sslmode=disable" -path migrations up 1

migrate_down:
	migrate -database "postgres://postgres:@localhost:5432/go_service?sslmode=disable" -path migrations down 1

migrate_create:
	 migrate create -ext sql -dir migrations/ ${name}
