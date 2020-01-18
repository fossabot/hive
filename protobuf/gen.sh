#!/usr/bin/env bash

GO_HIVE="../go/hive"
JS_HIVE="../js/hive"

protoc \
-I=. \
-I=$GOPATH/src/ \
-I=$GOPATH/src/github.com/gogo/protobuf/protobuf \
--gogofaster_out=plugins=grpc\
:$GO_HIVE *.proto \
--plugin=protoc-gen-ts=../test/web/node_modules/.bin/protoc-gen-ts \
--ts_out=service=grpc-web:$JS_HIVE \
--js_out=import_style=commonjs,binary:${JS_HIVE} \

for f in $JS_HIVE/* ;
    do {
        echo "/* eslint-disable */ $(cat $f)" > $f;
        sed -i '' '/gogoproto_gogo_pb/d' $f;
    };
done
