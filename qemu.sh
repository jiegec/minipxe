#!/bin/sh
qemu-system-x86_64 -smp 4 -m 8192 -boot n -net nic -net tap,ifname=minipxe-test,script=setup-tap.sh,downscript=no -serial mon:stdio -nographic