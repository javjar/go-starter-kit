#!/usr/bin/env bash

GO_WORK_DIR=${GO_WORK_DIR:-$GOPATH/src}
cd ${GO_WORK_DIR}
# export GO111MODULE=on
exec "$@"
