#!/bin/sh

chain33_path=$(go list -f '{{.Dir}}' "github.com/PhenixChain/devchain")
protoc --go_out=plugins=grpc:../types ./*.proto --proto_path=. --proto_path="${chain33_path}/chain33/types/proto/"
