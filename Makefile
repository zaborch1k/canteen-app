all: build gen-swag-docs

.PHONY: build
build:
	go build -o canteen-app cmd/http-server/main.go

.PHONY: gen-swag-docs
gen-swag-docs:
	swag init --dir ./cmd/http-server,./internal/adapter/http/api -o cmd/docs

.PHONY: run
run:
	./canteen-app