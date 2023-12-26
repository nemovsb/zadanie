OS=$(shell uname -o)
PROJECTNAME=zadanie

.PHONY: build-for-compose

# build: build for local OS
build:
	@echo
	go build -o ./build/bin/zadanie -ldflags "-s -w" zadanie/
	@echo ">complete"
	@echo

## build for docker-compose
GOOS=linux
GOARCH=amd64
CGO_ENABLED=0
BUILDVARS=GOOS=${GOOS} GOARCH=${GOARCH} CGO_ENABLED=${CGO_ENABLED}
DOCKER_BUILDVARS=GOOS=linux GOARCH=${GOARCH} CGO_ENABLED=${CGO_ENABLED}

# build-for-alpine: build for docker alpine
build-for-alpine:
	@echo
	${DOCKER_BUILDVARS} go build -buildvcs -o ./build/bin/zadanie -ldflags "-s -w" zadanie/
	@echo ">build for alpine complete"
	@echo

# test: run all tests
test:
	@echo
	@echo ">Running unit tests..."
	go test -v ./...
	@echo ">All tests passed"
	@echo

COMPOSE_FILE=docker-compose.yml
COMPOSE_CMD=docker-compose --project-name ${PROJECTNAME} --file ${COMPOSE_FILE}
COMPOSE_PULL_CMD=${COMPOSE_CMD} pull

## compose-up: raise the whole project from docker-compose.yml 
.PHONY: compose-up
compose-up: test build-for-alpine
	@echo "> Raise the whole project from docker-compose.yml..."
	${COMPOSEPULL_CMD}
	${COMPOSE_CMD} up --build --detach
	@echo "> Project raised"

## compose-down: destroy everything raised from docker-compose.yml
.PHONY: compose-down
compose-down:
	@echo "> Destroying everything raised from docker-compose.local.yml..."
	${COMPOSE_LOCAL_CMD} down --remove-orphans
	@echo "> Everything destroyed"
