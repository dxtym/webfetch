build:
	@/snap/bin/go build -o tmp/main cmd/main.go

run:
	@/snap/bin/go run cmd/main.go

.PHONY: build run
