#!/bin/bash

PROTOC_CMD_DIR=/Users/miksmolin/go/bin/protoc-gen-go
PROTOC_GEN_GRPC_GATEWAY_CMD_DIR=/Users/miksmolin/go/bin/protoc-gen-grpc-gateway
PROTOC_GEN_SWAGGER_CMD_DIR=/Users/miksmolin/go/bin/protoc-gen-swagger
PROJECT_DIR=/Users/miksmolin/ds-crm

echo "Start"
protoc --plugin=protoc-gen-go=$PROTOC_CMD_DIR -I$PROJECT_DIR/vendor.pb/ \
    	-I$PROJECT_DIR \
    	--go_out=plugins=grpc:$PROJECT_DIR/pkg \
      	$PROJECT_DIR/api/ds-crm.proto
protoc --plugin=protoc-gen-grpc-gateway=$PROTOC_GEN_GRPC_GATEWAY_CMD_DIR -I$PROJECT_DIR/vendor.pb/ \
    	-I$PROJECT_DIR \
    	--grpc-gateway_out=logtostderr=true:$PROJECT_DIR/pkg \
    	$PROJECT_DIR/api/ds-crm.proto
protoc --plugin=protoc-gen-swagger=$PROTOC_GEN_SWAGGER_CMD_DIR -I$PROJECT_DIR \
    	-I$PROJECT_DIR/vendor.pb/ \
    	--swagger_out=logtostderr=true:$PROJECT_DIR/pkg \
    	api/ds-crm.proto
