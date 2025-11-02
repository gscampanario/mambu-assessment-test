stubs:
	@echo "Generating stubs..."
	mockgen -source=./infrastructure/db/transaction_repository.go -destination=./mock/transaction_repository_mock.go -package=mock
	@echo "Stubs generated"

tests:
	@echo "Running tests..."
	make stubs
	go test -v ./...
	@echo "Tests completed"

run:
	@echo "Starting application..."
	CGO_ENABLED=1 go run ./main.go