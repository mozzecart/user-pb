VERSION ?= $(shell git describe --tags 2>/dev/null | cut -c 2-)

.PHONY: proto
proto:
	protoc --go_opt= --go_out=plugins=grpc:. *.proto

