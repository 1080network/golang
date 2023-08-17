#!/bin/bash

set -x
SDK=$1
INPUT_DIR=$SDK/proto
OUTPUT_DIR=.

VENDOR=$PWD/vendor

pushd $INPUT_DIR
echo "GENERATING GOLANG proto bindings"
#protoc --go_out $OUTPUT_DIR --go_opt paths=import \
#      --go-grpc_out $OUTPUT_DIR --go-grpc_opt paths=import \
#      --validate_out="lang=go:$OUTPUT_DIR" \
#      -I . \
#      -I $VENDOR/github.com/envoyproxy/protoc-gen-validate \
#      $(find . -name \*.proto)
protoc --go_out $OUTPUT_DIR --go_opt paths=import \
      --go-grpc_out $OUTPUT_DIR --go-grpc_opt paths=import \
      -I . \
      -I $VENDOR/github.com/envoyproxy/protoc-gen-validate \
      $(find . -name \*.proto)

if [[ $? -ne 0 ]] ; then
	popd
	echo "Miserable failure on your part man"
	exit 1
fi

GITHUB_MODULE_ADDRESS_SED="github.com\\/1080network\\/golang"
IMPORT_PATH="\\\"$GITHUB_MODULE_ADDRESS_SED\\/$SDK\\/proto"
for gofile in `find . -name "*.pb.go"` ; do
	sed -i.bak "s/\"mica\\//${IMPORT_PATH}\\/mica\\//g" $gofile
      sed -i.bak "s/\"micashared\\//${IMPORT_PATH}\\/micashared\\//g" $gofile
done
#clean up all the backup files
for file in `find . -name "*.bak"` ; do
	rm $file
done

popd
echo "Successful proto generation"
