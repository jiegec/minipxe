#!/bin/sh
dd if=/dev/zero of=pflash0.img bs=1048576 count=32
dd if=/dev/zero of=pflash1.img bs=1048576 count=32
dd if=/usr/share/qemu-efi-riscv64/RISCV_VIRT_CODE.fd of=pflash0.img conv=notrunc
dd if=/usr/share/qemu-efi-riscv64/RISCV_VIRT_VARS.fd of=pflash1.img conv=notrunc
qemu-system-riscv64 -machine virt,pflash0=pflash0,pflash1=pflash1,acpi=off -smp 2 -m 8192 -blockdev node-name=pflash0,driver=file,read-only=on,filename=pflash0.img -blockdev node-name=pflash1,driver=file,filename=pflash1.img -boot n -net nic -net tap,ifname=minipxe-test,script=setup-tap.sh,downscript=no -serial mon:stdio -device virtio-gpu-pci -device qemu-xhci -device usb-kbd -device usb-tablet
