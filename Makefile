build:
	@/snap/bin/go build -o tmp/main cmd/webfetch/main.go

run:
	@/snap/bin/go run cmd/webfetch/main.go

.PHONY: build run
