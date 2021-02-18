.PHONY: all
all: test benchmark

.PHONY: test
test:
	@go test -v ./...

.PHONY: benchmark
benchmark:
	@GOMAXPROCS=1 go test -v ./... -run=nil -bench=. -benchmem -benchtime=10000x

.PHONY: build
build:
	@go build -gcflags '-m -m' ./verhoeff.go