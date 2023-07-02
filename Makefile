build:
	@go build -o bin/azario

run: build
	@./bin/azario

test:
	@go test -v ./..