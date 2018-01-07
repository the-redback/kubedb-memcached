#!/bin/bash
set -xeou pipefail

GOPATH=$(go env GOPATH)
REPO_ROOT="$GOPATH/src/github.com/kubedb/memcached"

export APPSCODE_ENV=prod

pushd $REPO_ROOT

rm -rf dist

./hack/docker/mc-operator/setup.sh
./hack/docker/mc-operator/setup.sh release

rm dist/.tag

popd
