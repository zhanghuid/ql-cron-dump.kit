PROJECT=cron
OUTPUT=bin/$(PROJECT)
VERSION=0.0.1

.PHONY: build
build:
	go mod vendor
	CGO_ENABLED=0 go build -ldflags="-s" -mod=vendor -trimpath -o $(OUTPUT)-$(VERSION)-darwin

clean:
	go clean
	rm -f $(OUTPUT)
	rm -rf vendor/

run: build
	$(OUTPUT)

build-linux:
	go mod vendor
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s" -mod=vendor -trimpath -o $(OUTPUT)

# 精简版程序
build-linux-upx:
	go mod vendor
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s" -mod=vendor -trimpath -o $(OUTPUT)-dump
	upx -6 $(OUTPUT)-dump


build-mac:
	go mod vendor
	CGO_ENABLED=0 go build -ldflags="-s" -mod=vendor -trimpath -o $(OUTPUT)-$(VERSION)-darwin
