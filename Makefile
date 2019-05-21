BINARY = awsdyndns
GOARCH = amd64

VERSION=1
COMMIT=$(shell git rev-parse HEAD)
BRANCH=$(shell git rev-parse --abbrev-ref HEAD)
CURRENT_DIR=$(shell pwd)
BUILD_DIR=${CURRENT_DIR}/bin

# Setup the -ldflags option for go build here, interpolate the variable values
LDFLAGS = -ldflags "-X main.VERSION=${VERSION} -X main.COMMIT=${COMMIT} -X main.BRANCH=${BRANCH}"

# Build the project	
all: clean linux darwin windows

configure:
	mkdir ${BUILD_DIR}

linux: 
	GOOS=linux GOARCH=${GOARCH} go build ${LDFLAGS} -o ${BUILD_DIR}/${BINARY}-v${VERSION}-linux-${GOARCH} .
	# Build arm version for Raspberry Pi
	GOOS=linux GOARCH=arm GOARM=5 go build ${LDFLAGS} -o ${BUILD_DIR}/${BINARY}-v${VERSION}-linux-arm .

darwin:
	GOOS=darwin GOARCH=${GOARCH} go build ${LDFLAGS} -o ${BUILD_DIR}/${BINARY}-v${VERSION}-darwin-${GOARCH} .

windows:
	GOOS=windows GOARCH=${GOARCH} go build ${LDFLAGS} -o ${BUILD_DIR}/${BINARY}-v${VERSION}-windows-${GOARCH}.exe .

fmt:
	go fmt $$(go list ./... | grep -v /vendor/)

clean:
	rm -rf ${BUILD_DIR}

.PHONY: linux darwin windows fmt clean