#!/usr/bin/env bash

set -eoux pipefail

REPO_NAME=memcached
OPERATOR_NAME=mc-operator

# get concourse-common
pushd $REPO_NAME
git status
git subtree pull --prefix hack/concourse/common https://github.com/kubedb/concourse-common.git master --squash -m 'concourse'
popd

source $REPO_NAME/hack/concourse/common/init.sh

pushd "$GOPATH"/src/github.com/kubedb/$REPO_NAME

# build and push docker-image
./hack/builddeps.sh
export APPSCODE_ENV=dev
export DOCKER_REGISTRY=kubedbci

./hack/docker/$OPERATOR_NAME/make.sh build
./hack/docker/$OPERATOR_NAME/make.sh push

pushd $GOPATH/src/github.com/kubedb/$REPO_NAME

# run tests
source ./hack/deploy/setup.sh --docker-registry=kubedbci
./hack/make.py test e2e --v=1 --selfhosted-operator=true --ginkgo.flakeAttempts=2
