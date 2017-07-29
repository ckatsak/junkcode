#!/bin/bash
#
# Checks whether a route exists on the system, which forwards the traffic
# through a specified interface


mytest() {
	#local dev="docker0"
	local dev="focker0"

	local present=$(ip route ls | awk -v tapif=$dev '{for(i=0;i<NF;i++){if($i=="dev"){if($(i+1)==tapif){print 1; exit; } next; }}}')

	echo -e "\$present == $present\n"

	#if [ -n "$present" ]; then
	if [ -z "$present" ]; then
		echo "ERROR"
	fi
}

mytest

#local absent=$($IP route ls | awk -v tapif=$tapif '{ for (i=0; i<NF; i++){ if ($i == "dev") {\
	#	if ($(i+1) == tapif) { print 1; exit; } next; } } }')
## if it's present, then $present == 1, else $present == 0
#if [ $absent ]; then
#	echo "ERROR: Tap interface '$tapif' already exists!"
#	cleanup
#	exit 8
#fi  

