


all: 
	GOOS=linux \
	GOARCH=arm \
	GOARM=5 \
	CC=/pitools/arm-bcm2708/gcc-linaro-arm-linux-gnueabihf-raspbian-x64/bin/arm-linux-gnueabihf-gcc \
	CGO_ENABLED=1 \
	go build -x