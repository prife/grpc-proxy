#!/bin/bash

function protoc_gen {
   protoc --go_out=. --go_opt=paths=source_relative \
       --go-grpc_out=. --go-grpc_opt=paths=source_relative $@
}

protoc_gen perfdog/fakeperfdog/perfdog.proto
protoc_gen perfdog/perfdog/perfdog.proto
