
build:
	go build -o bin/cli ./cmd/cli

run:
	go run ./cmd/cli

test:
	go test ./...

clean:
	rm -rf bin
