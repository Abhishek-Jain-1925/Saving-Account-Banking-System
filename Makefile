run: ## Run e-commerce poject on host machine
	go run cmd/main.go

clean:
	clear
	
cleanDB: ## Clean database file for a fresh start
	rm repository/bank.db

test: ## Run all unit tests in the project
	go test -v ./...

test-cover: ## Run all unit tests in the project with test coverage
	go test -v ./... -covermode=count -coverprofile=coverage.out