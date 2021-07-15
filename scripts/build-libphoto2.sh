#!/bin/bash -ex
#
# Usage ./build-libgphoto2
#

BUILD_PATH=/tmp/build
INSTALL_PATH=/opt/shot-capture

LIBGPHOTO2_VERSION_MAJOR=2
LIBGPHOTO2_VERSION_MINOR=5
LIBGPHOTO2_VERSION_REVISION=27

LIBGPHOTO2_VERSION=$LIBGPHOTO2_VERSION_MAJOR.$LIBGPHOTO2_VERSION_MINOR.$LIBGPHOTO2_VERSION_REVISION

sudo apt-get install -y autoconf automake gcc pkg-config libtool gettext make libusb-1.0-0-dev libxml2-dev libjpeg-dev libexif-dev libcurl4-openssl-dev

sudo mkdir -p $INSTALL_PATH

mkdir -p $BUILD_PATH

pushd .
cd $BUILD_PATH

if [ ! -f libgphoto2-$LIBGPHOTO2_VERSION.tar.gz ]; then
    curl -LO https://github.com/gphoto/libgphoto2/releases/download/v$LIBGPHOTO2_VERSION/libgphoto2-$LIBGPHOTO2_VERSION.tar.gz
fi

if [ -d libgphoto2-$LIBGPHOTO2_VERSION ]; then
    rm -rf libgphoto2-$LIBGPHOTO2_VERSION
fi

tar xzf libgphoto2-$LIBGPHOTO2_VERSION.tar.gz

cd libgphoto2-$LIBGPHOTO2_VERSION

./configure --prefix=$INSTALL_PATH

make

sudo make install

popd

exit 0
