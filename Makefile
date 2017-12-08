PKG_OS = $$(if [ -z ${GOOS} ] ; then echo darwin ; else echo ${GOOS} ; fi)
PKG_ARCH = $$(if [ -z ${GOARCH} ] ; then echo amd64 ; else echo ${GOARCH} ; fi)

.PHONY: all
all: clean build benchmark

.PHONY: clean
clean:
	@go clean

.PHONY: build
build:
	@GOOS=$(PKG_OS) GOARCH=$(PKG_ARCH) go build -a -o pkg/verhoeff-$(PKG_OS)-$(PKG_ARCH) src/verhoeff.go src/main.go

.PHONY: test
test:
	@go test -v ./src/...

.PHONY: benchmark
benchmark:
	@go test -v ./src/... -bench=.
