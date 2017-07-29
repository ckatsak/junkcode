#!/bin/bash
#
# test notify-send

sleep 1 &
wait $!

notify-send "$0:" \
	"All VMs look dead now.\nI should clean up the host network settings.\nPlease insert password for sudo"
sleep 1
exit 0
