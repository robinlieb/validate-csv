# ==============================================================================
# Build

build:
	go build ./cmd/main.go

clean:
	rm main

# ==============================================================================
# Running tests within the local computer

test:
	go test ./... -count=1