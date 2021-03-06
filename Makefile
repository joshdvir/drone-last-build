
clean:
	go clean -i ./...

deps:
	go get -t ./...

build:
	go build

windows:
	docker run --rm -v /Users/shuky/.go/src/drone-last-build:/go/src/drone-last-build -w /go/src/drone-last-build -e GOOS=windows -e GOARCH=amd64 golang:1.6 go get && go build -v -o build/drone-last-build-windows-amd64
	hub release create -a build/drone-last-build-windows-amd64 -f build/drone-last-build-windows-amd64 drone-last-build

linux:
	docker run --rm -v /Users/shuky/.go/src/drone-last-build:/go/src/drone-last-build -w /go/src/drone-last-build golang:1.6 go get && go build -v -v -o build/drone-last-build-linux-amd64
	hub release create -a build/drone-last-build-linux-amd64 -f build/drone-last-build-linux-amd64 drone-last-build

darwin:
	docker run --rm -v /Users/shuky/.go/src/drone-last-build:/go/src/drone-last-build -w /go/src/drone-last-build -e GOOS=darwin -e GOARCH=amd64 golang:1.6 go get && go build -v -v -o build/drone-last-build-darwin-amd64
	hub release create -a build/drone-last-build-darwin-amd64 -f build/drone-last-build-darwin-amd64 drone-last-build

all: darwin linux windows