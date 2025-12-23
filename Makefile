all: build

.PHONY: build
build:
	go build -o canteen-app cmd/http-server/main.go
