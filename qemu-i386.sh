#!/bin/sh
qemu-system-i386 -accel kvm -smp 4 -m 8192 -boot n -net nic -net tap,ifname=minipxe-test,script=setup-tap.sh,downscript=no -serial mon:stdio
