GOCMD=go
GOTEST=$(GOCMD) test
GOVET=$(GOCMD) vet
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
VERSION?=0.0.1

.PHONY: all

GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
RESET  := $(shell tput -Txterm sgr0)

all: up

up:
	docker-compose up --build -d

down:
	docker-compose down

clean:
	$(GOCLEAN)
	rm -f bin/*

build: clean
	$(GOBUILD) -o bin/ ./pkg/...

get:
	$(GO) get ./...
	$(GO) mod verify
	$(GO) mod tidy

update:
	$(GO) get -u -v all
	$(GO) mod verify
	$(GO) mod tidy

lint:
	golangci-lint run

test:
	$(GOTEST) -v -race ./...

coverage:
	$(GOTEST) -cover \
		-covermode=count \
		-coverprofile=profile.cov ./...
	$(GOCMD) tool cover \
		-func profile.cov


help:
	@echo 'Usage:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo 'Targets:'
	@echo '  ${YELLOW}up${RESET}           ${GREEN}Build the docker containers and start it${RESET}'
	@echo '  ${YELLOW}down${RESET}         ${GREEN}Finish the docker containers${RESET}'
	@echo '  ${YELLOW}clean${RESET}        ${GREEN}Clean /bin directory${RESET}'
	@echo '  ${YELLOW}build${RESET}        ${GREEN}Build Go code. Required Golang to be installed${RESET}'
	@echo '  ${YELLOW}get${RESET}          ${GREEN}Get Golang packages${RESET}'
	@echo '  ${YELLOW}update${RESET}       ${GREEN}Update Golang packages${RESET}'
	@echo '  ${YELLOW}lint${RESET}         ${GREEN}Lint Go code${RESET}'
	@echo '  ${YELLOW}test${RESET}         ${GREEN}Test Go code${RESET}'
	@echo '  ${YELLOW}coverage${RESET}     ${GREEN}Export Go tests coverage${RESET}'
	@echo ''
