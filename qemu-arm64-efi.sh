#!/bin/sh
dd if=/dev/zero of=pflash0.img bs=1048576 count=64
dd if=/dev/zero of=pflash1.img bs=1048576 count=64
dd if=/usr/share/qemu-efi-aarch64/QEMU_EFI.fd of=pflash0.img conv=notrunc
qemu-system-aarch64 -cpu max -machine virt -smp 4 -m 8192 -drive if=pflash,file=pflash0.img,format=raw,readonly=on -drive if=pflash,file=pflash1.img,format=raw -boot n -net nic -net tap,ifname=minipxe-test,script=setup-tap.sh,downscript=no -serial mon:stdio -device virtio-gpu-pci -device qemu-xhci -device usb-kbd -device usb-tablet
