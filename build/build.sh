#!/bin/sh

# execute by Dockerfile, work dir is WORKDIR "/app"

export GOPROXY=https://goproxy.cn,direct
export GO111MODULE=on

SCRIPT_DIR="."
PROJECT_DIR="${SCRIPT_DIR}/../"
RES_FILE_NAME="local-api-gateway"

# build_app 编译
build_app() {
    echo "enter download_dependency"
    if [ ! -d "${PROJECT_DIR}" ]; then
        echo "do not have project dir"
        return 1
    fi

    go mod download
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static" -w -s' -o "./${RES_FILE_NAME}" ./src/main

    chmod +x ./${RES_FILE_NAME}
    return 0
}

# main
main() {
    build_app
    build_res=$?
    if [ "$build_res" != 0 ]; then
        echo "build error"
    fi
    echo "build success"
}

main "$@"