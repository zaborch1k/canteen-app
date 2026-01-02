all: build gen-swag-docs

.PHONY: build
build:
	go build -o canteen-app cmd/http-server/main.go

.PHONY: gen-swag-docs
gen-swag-docs:
	swag init --dir ./cmd/http-server,./internal/adapter/http/api -o cmd/docs

.PHONY: gen-mocks
gen-mocks:
	mockery

.PHONY: run
run:
	./canteen-app

.PHONY: test-auth-api
test-auth-api: 
	go test -v canteen-app/internal/adapter/http/api/