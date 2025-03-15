build:
	go build -o backhome ./cmd/backhome

install:
	go install ./cmd/backhome

lint:
	golangci-lint run ./...

clear:
	rm -rf ./local/
	rm -rf ./local.backhome/