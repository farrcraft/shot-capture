#!/bin/bash -ex

GOLANG_VERSION=1.16.6

# pi is an ARMv7 so closest working go binaries are "v6l"
PLATFORM=amd64

pushd .

cd /tmp/

curl -LO https://golang.org/dl/go$GOLANG_VERSION.linux-$PLATFORM.tar.gz

if [ -d /usr/local/go ]; then
    sudo rm -rf /usr/local/go
fi

sudo tar -C /usr/local -xzf go$GOLANG_VERSION.linux-$PLATFORM.tar.gz

popd

exit 0
