clean:
	rm -f main

build: clean
	go build cmd/main/main.go

run: build
	./main

test:
	go test -v -count=1 ./...

.PHONY: cover
cover:
	go test -short -count=1 -race -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out
	rm coverage.out