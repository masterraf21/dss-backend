.PHONY: mock test
include .env

# run: test-api test-service
run:
	./cmds/env .env go run main.go

build:
	./cmds/env .env go build main.go

dev: 
	./cmds/env .env
	gin --appPort ${PORT} --all -i main.go

test:
	go get -u github.com/kyoh86/richgo
	./cmds/env env-test richgo test -count=1 ./... -v -cover
	go mod tidy

test-repo:
	go get -u github.com/kyoh86/richgo
	./cmds/env env-test richgo test -count=1 ./repositories/mongodb -v -cover
	go mod tidy

test-usecase:
	go get -u github.com/kyoh86/richgo
	./cmds/env env-test richgo test -count=1 ./usecases -v -cover
	go mod tidy

tidy:
	go mod tidy

download:
	go mod download

docker-dev-up:
	docker-compose -f docker-compose.dev.yml up -d
docker-dev-down:
	docker-compose -f docker-compose.dev.yml down
docker-dev-stop:
	docker-compose -f docker-compose.dev.yml stop

mock:
	@mockery --dir models --all