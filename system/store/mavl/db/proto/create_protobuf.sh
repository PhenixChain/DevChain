#!/bin/sh

chain33_path=$(go list -f '{{.Dir}}' "github.com/PhenixChain/devchain")
protoc --go_out=plugins=grpc:../ticket ./*.proto --proto_path=. --proto_path="${chain33_path}/types/proto/"
