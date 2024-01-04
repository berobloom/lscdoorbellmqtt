export PACKAGE_NAME := lscdoorbellmqtt
export GOOS := linux
export GOARCH := arm
export GOARM := 7
export GOMIPS := softfloat
export CC := $(TOOLCHAIN_BIN_DIR)/arm-openipc-linux-musleabi-gcc
export STRIP := $(TOOLCHAIN_BIN_DIR)/arm-openipc-linux-musleabi-strip
export CGO_ENABLED := 1

.PHONY: all clean

all: build

build:
	go build -o $(PACKAGE_NAME)
	$(STRIP) $(PACKAGE_NAME)
	mkdir -p output/$(PACKAGE_NAME)
	mv $(PACKAGE_NAME) output/$(PACKAGE_NAME)/
	cp sound/boot.wav sound/dingdong.wav output/$(PACKAGE_NAME)/

clean:
	rm -rf $(PACKAGE_NAME)/
