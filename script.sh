#!/bin/bash
BUILD_PATH=bin
SCRIPT_NAME=script.sh
BUILD_ACTION=build

Kiểm tra tham số đầu vào
if [ "$1" == ${BUILD_ACTION} ]; then
    CGO_ENABLED=0 go build -v -o ${BUILD_PATH}/ . && echo ">>> Built successfully"
elif [ "$1" == "run" ]; then
    sh ${SCRIPT_NAME} ${BUILD_ACTION}
    ${BUILD_PATH}
elif [ "$1" == "clean" ]; then
    go clean
    rm -f ${BUILD_PATH}
elif [ "$1" == "test" ]; then
    go test ./...
fi
exit 0