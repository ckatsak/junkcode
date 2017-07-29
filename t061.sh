#!/bin/bash
#
# test if route to interface exists	: comment in and out mocker0

VAR="docker0"
#VAR="mocker0"

#ip route ls | awk -v var="$VAR" '\
#{\
#	for (i=0; i<NF; i++) {\
#		if ($i == "dev") {\
#			if ($(i+1) == var) {\
#				exit(42);\
#			}\
#			next;\
#		}\
#	}\
#}'
present=$(ip route ls | awk -v br=$VAR '{ for (i=0; i<NF; i++) { if ($i == "dev") {\
	if ($(i+1) == br) { print 1; exit; } next; } } }')
echo "\$present=$present"
if [ ! $present ]; then
	echo "NOT present!"
fi
