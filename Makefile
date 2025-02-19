build:
	@/snap/bin/go build -o bin/main cmd/main.go

run:
	@/snap/bin/go run cmd/main.go

.PHONY: build run
