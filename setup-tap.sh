#!/bin/sh
ip link set "$1" up
ip addr add 192.168.0.1/24 dev "$1"
