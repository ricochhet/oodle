CUSTOM=-X 'main.buildDate=$(shell date)' -X 'main.gitHash=$(shell git rev-parse --short HEAD)' -X 'main.buildOn=$(shell go version)'
LDFLAGS=$(CUSTOM) -w -s -extldflags=-static

GO_BUILD=go build -trimpath -ldflags "$(LDFLAGS)"
BUILD_OUTPUT=build
ASSET_PATH=assets

APP_PATH=./cmd/oodle

.PHONY: all
all: oodle-linux oodle-linux-arm64 oodle-darwin oodle-darwin-arm64 oodle-windows

.PHONY: all-windows
all-windows: oodle-windows

.PHONY: fmt
fmt:
	gofumpt -l -w -extra .

.PHONY: tidy
tidy:
#	go get -u ./...
	@echo "[main] tidy"
	go mod tidy

.PHONY: update
update:
	@echo "[main] tidy"
	go get -u ./...

.PHONY: lint
lint: fmt
# golangci-lint cache clean
	@echo "[main] golangci-lint"
	golangci-lint run ./... --fix

.PHONY: test
test:
	go test ./...

.PHONY: deadcode
deadcode:
	deadcode ./...

.PHONY: syso
syso:
	windres $(APP_PATH)/app.rc -O coff -o $(APP_PATH)/app.syso

.PHONY: png-to-icos
png-to-icos:
	magick $(ASSET_PATH)/win-icon.png -background none -define icon:auto-resize=256,128,64,48,32,16 $(ASSET_PATH)/win-icon.ico

.PHONY: copy-libs
copy-libs:
#	cp -r libs/liboo2corelinux64.so.9 build/liboo2corelinux64.so.9
#	cp -r libs/oo2core_9_win64.dll build/oo2core_9_win64.dll

.PHONY: oodle-linux
oodle-linux: copy-libs fmt
	CGO_ENABLED=1 GOOS=linux GOARCH=amd64 $(GO_BUILD) -o $(BUILD_OUTPUT)/oodle-linux $(APP_PATH)

.PHONY: oodle-linux-arm64
oodle-linux-arm64: copy-libs fmt
	CGO_ENABLED=1 GOOS=linux GOARCH=arm64 $(GO_BUILD) -o $(BUILD_OUTPUT)/oodle-linux-arm64 $(APP_PATH)

.PHONY: oodle-darwin
oodle-darwin: copy-libs fmt
	CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 $(GO_BUILD) -o $(BUILD_OUTPUT)/oodle-darwin $(APP_PATH)

.PHONY: oodle-darwin-arm64
oodle-darwin-arm64: copy-libs fmt
	CGO_ENABLED=1 GOOS=darwin GOARCH=arm64 $(GO_BUILD) -o $(BUILD_OUTPUT)/oodle-darwin-arm64 $(APP_PATH)

.PHONY: oodle-windows
oodle-windows: copy-libs fmt
	CGO_ENABLED=1 GOOS=windows GOARCH=amd64 $(GO_BUILD) -o $(BUILD_OUTPUT)/oodle.exe $(APP_PATH)