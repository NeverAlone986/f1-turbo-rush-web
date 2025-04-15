.PHONY: build run test clean

build:
	go build -o bin/f1-turbo-rush ./cmd/web

run: build
	./bin/f1-turbo-rush

test:
	go test -v ./...

clean:
	rm -rf bin/
