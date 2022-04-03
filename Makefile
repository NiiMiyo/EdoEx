VERSION      = 1.0
OUTPUT       = bin
CROSS_OUTPUT = build
BUILDFILES   = buildfiles

clean:
	go clean
	rm ${OUTPUT}/* -rf

build: clean
	GOOS=$(shell go env GOHOSTOS) GOARCH=$(shell go env GOHOSTARCH) go build -o ${OUTPUT}/ .
	cp ${BUILDFILES}/* ${OUTPUT}/ -r


cross-clean:
	go clean
	rm ${CROSS_OUTPUT}/* -rf


cross-build: cross-clean
	GOOS=windows GOARCH=amd64 go build -o ${CROSS_OUTPUT}/win_x64-${VERSION}/ .
	cp ${BUILDFILES}/* ${CROSS_OUTPUT}/win_x64-${VERSION}/ -r

	GOOS=windows GOARCH=386 go build -o ${CROSS_OUTPUT}/win_x32-${VERSION}/ .
	cp ${BUILDFILES}/* ${CROSS_OUTPUT}/win_x32-${VERSION}/ -r

	GOOS=linux GOARCH=amd64 go build -o ${CROSS_OUTPUT}/linux_x64-${VERSION}/ .
	cp ${BUILDFILES}/* ${CROSS_OUTPUT}/linux_x64-${VERSION}/ -r

	GOOS=linux GOARCH=386 go build -o ${CROSS_OUTPUT}/linux_x32-${VERSION}/ .
	cp ${BUILDFILES}/* ${CROSS_OUTPUT}/linux_x32-${VERSION}/ -r

	GOOS=darwin GOARCH=amd64 go build -o ${CROSS_OUTPUT}/darwin_x64-${VERSION}/ .
	cp ${BUILDFILES}/* ${CROSS_OUTPUT}/darwin_x64-${VERSION}/ -r
