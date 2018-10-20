build:
	go build -o bin/locker ./cmd/locker/

clean:
	docker rmi locker
	rm -rf vendor/*
	rm -rf bin/*

deps:
	go get -u github.com/golang/dep/cmd/dep
	dep ensure

docker-build:
	docker build -t locker .

docker-test: 
	docker run golang:latest cd /go/src/github.com/ryanhartje/locker && go test ./pkg/locker ./cmd/locker

test:
	go test -v ./pkg/locker
	go test -v ./cmd/locker
