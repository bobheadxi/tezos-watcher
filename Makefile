TEZOS_CLIENT = $(HOME)/tezos/tezos-client

all: deps cli

.PHONY: deps
deps:
	dep ensure -v

.PHONY: cli
cli:
	go install ./cmd/tezos-watcher
