deps:
	go get -u github.com/golang/dep/cmd/dep
	dep ensure

test:
	go test -v ./pkg/locker
	go test -v ./cmd/locker

build:
	go build -o bin/locker ./cmd/locker/
