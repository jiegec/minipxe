# minipxe

DHCP and TFTP server in one binary.

Boot process:

1. Firmware enters PXE boot mode
2. Firmware gets IP address and TFTP address of iPXE from DHCP server
3. Firmware downloads iPXE from TFTP server and runs it
4. iPXE gets IP address and TFTP address of iPXE config from DHCP server
5. iPXE executes iPXE config and boots live cd over network

## Build ipxe

Run `build.sh` in `root-ipxe` directory.

## Setup up PXE environment using isc-dhcp-server & tftpd-hpa

minipxe serves as a simple alternative to isc-dhcp-server + tftpd-hpa combination. But here also gives instructions on how to setup PXE environment based on isc-dhcp-server and tftpd-hpa:

1. Install isc-dhcp-server & tftpd-hpa
2. Add network interface to INTERFACESv4 in `/etc/default/isc-dhcp-server`
3. Append dhcp config to `/etc/dhcp/dhcpd.conf` (substitute IP addresses if necessary):

```
option client-architecture code 93 = unsigned integer 16;
subnet 192.168.1.0 netmask 255.255.255.0 {
  interface eno1;
  option routers 192.168.1.1;
  option domain-name-servers 114.114.114.114;

  range 192.168.1.100 192.168.1.200;
  next-server 192.168.1.1;

  if exists user-class and option user-class = "iPXE" {
    option vendor-class-identifier "PXEClient";
    filename "tftp://192.168.1.1/ipxe.cfg";
  } elsif option client-architecture = encode-int(0, 16) {
    filename "bin-i386-pcbios/undionly.kpxe";
  } elsif option client-architecture = encode-int(7, 16) {
    filename "bin-x86_64-efi/ipxe.efi";
  } elsif option client-architecture = encode-int(11, 16) {
    filename "bin-arm64-efi/ipxe.efi";
  }
}
```

4. Build iPXE:

```shell
git clone https://github.com/ipxe/ipxe.git
cd ipxe/src
make bin-i386-pcbios/undionly.kpxe NO_WERROR=1 -j32
make bin-x86_64-efi/ipxe.efi NO_WERROR=1 -j32
make bin-arm64-efi/ipxe.efi NO_WERROR=1 CROSS_COMPILE=aarch64-linux-gnu- -j32
```

5. Place iPXE and iPXE config under `/srv/tftp`
6. Start tftpd-hpa, change `/etc/default/tftpd-hpa` if necessary
