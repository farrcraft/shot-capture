#!/bin/bash -ex

GOLANG_VERSION=1.16.5

pushd .

cd /tmp/

curl -LO https://golang.org/dl/go$GOLANG_VERSION.linux-amd64.tar.gz

if [ -d /usr/local/go ]; then
    sudo rm -rf /usr/local/go
fi

sudo tar -C /usr/local -xzf go$GOLANG_VERSION.linux-amd64.tar.gz

popd

exit 0
