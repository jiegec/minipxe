#!/bin/sh
cd ipxe/src
make bin-i386-pcbios/undionly.kpxe NO_WERROR=1 -j32
make bin-x86_64-efi/ipxe.efi NO_WERROR=1 -j32
