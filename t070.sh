#!/bin/bash
#
# Run:
#
#	$ ./t70.sh <interface_name>
#
# Finds out if a network interface is present on the system.

IP=$(which "ip")
AWK=$(which "awk")

iface_present() {
	#echo "  I'm in, \$1==$1"
	local ifaces=$($IP a s | $AWK '/^[1-9]/{print $2}' | $AWK -F ":" '{print $1}')
	for iface in $ifaces; do
		if [[ "$iface" == "$1" ]]; then
			echo 0
			return
		fi
	done
}

mytest() {
	local tapif=$1
	#iface_present $tapif
	#local present=$?
	##echo $present
	#if [ "$present" -eq 1 ]; then
	#	echo "NOT PRESENT"
	#else
	#	echo "PRESENT"
	#fi

	#echo; echo; echo;
	local present=$(iface_present $tapif)
	if [ -z "$present" ]; then
		echo "NOT PRESENT"
	else
		echo "PRESENT"
	fi
}

mytest $1
