#!/bin/bash
abs_dirname() {
    local cwd="$(pwd)"
    local path="$1"

    while [ -n "$path" ]; do
        cd "${path%/*}"
        local name="${path##*/}"
        path="$(readlink "$name" || true)"
    done

    pwd -P
    cd "$cwd"
}

build_osx() {
    GOOS=darwin GOARCH=amd64 go build
    mkdir -p bin/osx
    mv $cmdName bin/osx/
    echo Completed build for OSX.
}
build_linux() {
    GOOS=linux GOARCH=amd64 go build
    mkdir -p bin/linux
    mv $cmdName bin/linux/
    echo Completed build for Linux.
}
build_windows() {
    GOOS=windows GOARCH=amd64 go build
    mkdir -p bin/windows
    mv ${cmdName}.exe bin/windows/
    echo Completed build for Windows.
}

cd $(abs_dirname "$0")

cmdName=difup
build_osx
build_linux
build_windows
