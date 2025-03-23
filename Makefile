ci: vet lint test build

vet:
	go vet ./...	

lint:
	golangci-lint run ./...

test:
	go test ./...

build:
	go build -o build/backhome ./cmd/backhome

install:
	go install ./cmd/backhome

clear:
	rm -rf ./local/
	rm -rf ./local.backhome/