#!/bin/sh
protoc --go_out=plugins=grpc:../grpc   *.proto
protoc --go_out=plugins=irpc:../http   *.proto


#protoDir="../proto"
#outDir="../proto"
#
#protoc -I ${protoDir}/ ${protoDir}/google/api/*.proto \
#--go_out ${outDir} \
#--go_opt paths=source_relative
#
#protoc -I ${protoDir}/ ${protoDir}/*.proto \
#--go_out ${outDir}/pb \
#--go_opt paths=source_relative \
#--go-grpc_out ${outDir}/pb \
#--go-grpc_opt paths=source_relative \
#--go-grpc_opt require_unimplemented_servers=false \
#--grpc-gateway_out ${outDir}/pb \
#--grpc-gateway_opt logtostderr=true \
#--grpc-gateway_opt paths=source_relative \
#--grpc-gateway_opt generate_unbound_methods=true \

#protoc -I ../../../ ../../../google/api/*.proto \
#--go_out  ../ggg \
#--go_opt paths=source_relative
#
#protoc    --go_out=plugins=irpc:. --go_opt paths=source_relative  --go-grpc_out=. --go-grpc_opt paths=source_relative   *.proto

#protoc -I ./ \
#	-I $GOPATH/src/google/api \
#	--go_out=plugins=irpc:. *.proto
