TEZOS_CLIENT = $(HOME)/tezos/tezos-client
TEZOS_RPC_PORT = 18731

all: deps cli

.PHONY: deps
deps:
	dep ensure -v

.PHONY: cli
cli:
	go install ./cmd/tezos-watcher

.PHONY: node
node:
	@echo Starting up Tezos sandbox node...
	@(cd test ; docker-compose up --build)

.PHONY: test
test:
	env TEZOS_RPC_PORT=$(TEZOS_RPC_PORT) go test -race -cover ./...
