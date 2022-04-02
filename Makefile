VERSION      = 1.0
OUTPUT       = bin
CROSS_OUTPUT = build

clean:
	go clean
	rm ${OUTPUT}/* -rf

build: clean
	GOOS=$(shell go env GOHOSTOS) GOARCH=$(shell go env GOHOSTARCH) go build -o ${OUTPUT}/ .


cross-clean:
	go clean
	rm ${CROSS_OUTPUT}/* -rf


cross-build: cross-clean
	GOOS=windows GOARCH=amd64 go build -o ${CROSS_OUTPUT}/win_x64-${VERSION}/ .

	GOOS=windows GOARCH=386 go build -o ${CROSS_OUTPUT}/win_x32-${VERSION}/ .

	GOOS=linux GOARCH=amd64 go build -o ${CROSS_OUTPUT}/linux_x64-${VERSION}/ .

	GOOS=linux GOARCH=386 go build -o ${CROSS_OUTPUT}/linux_x32-${VERSION}/ .

	GOOS=darwin GOARCH=amd64 go build -o ${CROSS_OUTPUT}/darwin_x64-${VERSION}/ .
