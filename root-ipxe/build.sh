#!/bin/sh
git submodule update --init --recursive
cd ipxe/src
make bin-i386-pcbios/undionly.kpxe NO_WERROR=1 -j32
make bin-x86_64-efi/ipxe.efi NO_WERROR=1 -j32
make bin-arm64-efi/ipxe.efi NO_WERROR=1 CROSS_COMPILE=aarch64-linux-gnu- -j32
