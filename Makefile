# go-set

clean:
	go clean

test:
	go test ./...

fmt:
	go fmt ./...

vet:
	go vet ./...

lint: fmt vet
	golangci-lint run ./...

