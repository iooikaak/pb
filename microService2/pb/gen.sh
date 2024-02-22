#!/bin/sh
protoc --go_out=plugins=grpc:../grpc   *.proto
protoc --go_out=plugins=irpc:../http   *.proto
