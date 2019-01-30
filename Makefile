.PHONY:
PROJECT      = pod-kicker
OUTFILE      = $(shell pwd)/$(PROJECT)
PK_TOKEN    ?=
PK_PORT     ?= 4200

.PHONY: test
test:
	go test ./...

.PHONY: build
build:
	go build

.PHONY: run
run:
	PK_TOKEN=$(PK_TOKEN) PK_PORT=$(PK_PORT) $(OUTFILE)
