NAME=pine
PKG=github.com/CDcoding2333/pet/cmd
VERSION=git-$(shell git describe --always --dirty)
IMAGE_TAG=$(VERSION)


linux:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
		go build -a -tags netgo -installsuffix netgo -installsuffix cgo -ldflags '-w -s' -ldflags "-X main.Version=$(VERSION)" \
		-o ./build/linux/pet $(PKG)
		upx ./build/linux/pet

darwin:
	GOOS=darwin GOARCH=amd64 \
		go build -a -tags netgo -installsuffix netgo -ldflags "-X main.Version=$(VERSION)" \
		-o ./build/darwin/pet $(PKG)