stubs:
	@echo "Generating stubs..."
	mockgen -source=./infrastructure/db/client_repository.go -destination=./mock/client_repository_mock.go -package=mock
	@echo "Stubs generated"

tests:
	@echo "Running tests..."
	make stubs
	go test -v ./...
	@echo "Tests completed"