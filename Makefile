# ==============================================================================
# Build

build:
	go build -o validate-csv ./cmd/

clean:
	rm validate-csv

# ==============================================================================
# Code analysis

vet:
	go vet ./...

lint:
	staticcheck -checks=all ./...

# ==============================================================================
# Running tests within the local computer

test:
	go test ./... -count=1