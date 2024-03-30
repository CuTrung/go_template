# !/bin/bash
PROJECT_NAME=go_template.exe
BUILD_PATH=bin/${PROJECT_NAME}
SCRIPT_NAME=script.sh

# Kiểm tra tham số đầu vào
if [ "$1" == "build" ]; then
    go build -o ${BUILD_PATH} main.go && echo ">>> Built successfully"
elif [ "$1" == "run" ]; then
    sh ${SCRIPT_NAME} build 
    ${BUILD_PATH}
elif [ "$1" == "clean" ]; then
    go clean
    rm -f ${BUILD_PATH}
fi
exit 0




