deps:
	go get -u github.com/golang/dep/cmd/dep
	dep ensure

docker-test: docker-build
	docker run locker go test ./pkg/locker ./cmd/locker

test:
	go test -v ./pkg/locker
	go test -v ./cmd/locker

build:
	go build -o bin/locker ./cmd/locker/

docker-build:
	docker build -t locker .

clean:
	docker rmi locker
	rm -rf vendor/*
	rm -rf bin/*