clean:
	rm -rf cmd/backhome/backhome-remote
	rm -rf backhome-local

build:
	go build -o backhome ./cmd/backhome
