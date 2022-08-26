# ==============================================================================
# Build

build:
	go build ./cmd/main.go

clean:
	rm main

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