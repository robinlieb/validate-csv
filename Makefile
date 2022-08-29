# ==============================================================================
# Build

build:
	go build -o validate-csv ./cmd/

clean:
	rm validate-csv

tidy:
	go mod tidy
	go mod vendor

deps-upgrade:
	go get -u -v ./...
	go mod tidy
	go mod vendor

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