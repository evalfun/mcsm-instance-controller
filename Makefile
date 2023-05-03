export PATH := $(GOPATH)/bin:$(PATH)
export GO111MODULE=on
LDFLAGS := -s -w


all: build-linux-amd64 build-windows-amd64 build-linux-arm64

fmt:
	go fmt ./...

fmt-more:
	gofumpt -l -w .


build-linux-amd64:
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -ldflags "$(LDFLAGS)" -o ./bin/mcsm-instance-controller-linux-amd64 ./cmd
	tar -czf ./release/mcsm-instance-controller-linux-amd64.tar.gz -C./ config.json -Cbin/ mcsm-instance-controller-linux-amd64 

build-linux-arm64:
	env CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -trimpath -ldflags "$(LDFLAGS)" -o ./bin/mcsm-instance-controller-linux-arm64 ./cmd
	tar -czf ./release/mcsm-instance-controller-linux-arm64.tar.gz -C./ config.json -Cbin/ mcsm-instance-controller-linux-arm64  

build-windows-amd64:
	env CGO_ENABLED=0 GOOS=windows GOARCH=amd64  go build -trimpath -ldflags "$(LDFLAGS)" -o ./bin/mcsm-instance-controller-windows-amd64.exe ./cmd
	tar -czf ./release/mcsm-instance-controller-windows-amd64.tar.gz -C./ config.json -Cbin/ mcsm-instance-controller-windows-amd64.exe 
	

clean:
	rm -f ./bin/mcsm-instance-controller-linux-amd64
	rm -f ./bin/mcsm-instance-controller-linux-arm64
	rm -f ./bin/mcsm-instance-controller-windows-amd64.exe
	rm -f ./release/mcsm-instance-controller-linux-amd64.tar.gz
	rm -f ./release/mcsm-instance-controller-linux-arm64.tar.gz
	rm -f ./release/mcsm-instance-controller-windows-amd64.tar.gz