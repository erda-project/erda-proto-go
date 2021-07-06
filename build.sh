#!/bin/bash

set -o errexit -o pipefail

# check parameters and print usage if need
usage() {
    echo "protoc.sh ACTION"
    echo "ACTION: "
    echo "    build        build proto to go files."
    echo "    clean        clean result files of protocol building."
    exit 1
}
if [ -z "$1" ]; then
    usage
fi

# build protocol
build_protocol() {
    MODULES=$(find ".proto" -type d);
    for path in ${MODULES}; do
        HAS_PROTO_FILE=$(eval echo $(bash -c "find "${path}" -maxdepth 1 -name *.proto 2>/dev/null" | wc -l));
        if [ ${HAS_PROTO_FILE} -gt 0 ]; then
            if [ -z "$(echo ${path#.proto})" ]; then
                continue; # skip .proto
            fi
            MODULE_PATH=${path#.proto/};
            echo "build module ${MODULE_PATH}";
            mkdir -p ${MODULE_PATH}/pb;
            mkdir -p ${MODULE_PATH}/client;
            gohub protoc protocol \
                 --include=.proto \
                 --msg_out="${MODULE_PATH}/pb" \
                 --service_out="${MODULE_PATH}/pb" \
                 --client_out="${MODULE_PATH}/client" \
                 --validate=true \
                 --json=true \
                 --json_opt=emit_defaults=true \
                 ${path}/*.proto
        fi;
    done;
    echo "";
    echo "build all proto successfully !";
}

# clean result files of building
clean_result() {
    MODULES=$(find ".proto" -type d);
    for path in ${MODULES}; do
        HAS_PROTO_FILE=$(eval echo $(bash -c "find "${path}" -maxdepth 1 -name *.proto 2>/dev/null" | wc -l));
        if [ ${HAS_PROTO_FILE} -gt 0 ]; then
            if [ -z "$(echo ${path#.proto})" ]; then
                continue; # skip .proto
            fi
            MODULE_PATH=${path#.proto/};
            rm -rf ${MODULE_PATH}
        fi;
    done;
}

case "$1" in
    "build")
        build_protocol
        ;;
    "clean")
        clean_result
        ;;
    *)
        usage
esac
