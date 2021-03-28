.PHONY: run test migrate migrate_test

run:
	go run cmd/app/main.go

migrate:
	migrate -database ${CAR_POSTGRES_URL} -path migrations up

migrate_test:
	migrate -database ${CAR_POSTGRES_TEST_URL} -path migrations up

test:
	go clean -testcache && go test ./... -v