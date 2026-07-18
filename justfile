# Install docs site dependencies
setup-docs:
	cd website && npm ci

# Build the project
build:
	go build ./...

# Run tests
test:
	go test -v ./...

# Serve docs locally with live reload
serve-docs:
	cd website && npm start

# Build static site output (without serving)
generate-docs:
	cd website && npm run build

# Run golangci-lint on the project
lint:
	golangci-lint run ./...

# Format all Go files in the project
fmt:
	gofmt -s -w .

# Vet doc code snippets and check region markers against the pages
check-docs:
	go vet -tags snippets ./website/snippets/
	./scripts/check-snippet-regions.sh

# Clean up local build artifacts
clean-docs:
	rm -rf website/build website/.docusaurus
