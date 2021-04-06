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
    PKG_PATH=$(go run github.com/erda-project/erda-infra/tools/gopkg github.com/erda-project/erda-infra/tools)
    MODULES=$(find ".proto" -type d);
    for path in ${MODULES}; do
        HAS_PROTO_FILE=$(eval echo $(bash -c "find "${path}" -maxdepth 1 -name *.proto 2>/dev/null" | wc -l));
        if [ ${HAS_PROTO_FILE} -gt 0 ]; then
            if [ -z "${path#.proto}" ]; then
                continue; # skip .proto
            fi
            echo "build module ${MODULE_PATH}";
            MODULE_PATH=${path#.proto/};
            PB_OUTPUT="./${MODULE_PATH}" "${PKG_PATH}/protoc.sh" protocol "${path}/*.proto"
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
            if [ -z "${path#.proto}" ]; then
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
