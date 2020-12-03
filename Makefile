APP_NAME=api

ENTRYPOINT="main.go"

## help: prints this message
.PHONY: help
help:
	@echo "Usage:"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

## clean: cleans the binary
.PHONY: clean
clean:
	@echo "Cleaning"
	@go clean

## tidy: clean unused dependencies
.PHONY: tidy
tidy:
	@echo "Tidying"
	@go mod tidy

## generate: generates go code (e.g. ent, swagger)
.PHONY: generate
generate: ent/ent.go docs/docs.go

# generate entity files
ent/ent.go: $(shell find ent/schema -type f)
	@echo "Generating: ent"
	@go generate ./ent

# generate swagger docs files
docs/docs.go: main.go $(shell find . -type f | grep controller.go)
	@export PATH=$(shell go env GOPATH)/bin:$$PATH; swag i

## test: runs tests recursively
.PHONY: test
test:
	@echo "Running tests"
	@go test -v ./...

## dev: runs without building
.PHONY: run
run:
	@echo "Running (dev)"
	@go run ${ENTRYPOINT}

## docker-build: builds and runs the docker image
.PHONY: docker
docker: 
	@echo "Running the docker image"
	@docker run \
		--rm \
		-it \
		--env-file ./.env \
		-e PORT=80 \
		-p 127.0.0.1:8081:80 \
		${APP_NAME}

## docker-build: builds the docker image
.PHONY: docker-build
docker-build: generate
	@echo "Building the docker image"
	docker build \
		--build-arg APP_NAME=${APP_NAME} \
		-t ${APP_NAME} \
		.
