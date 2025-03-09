clean:
	rm -rf cmd/backhome/backhome-remote
	rm -rf cmd/backhome/backhome-remote.backhome
	rm -rf backhome-local
	rm -rf backhome-local.backhome

build:
	go build -o backhome ./cmd/backhome
