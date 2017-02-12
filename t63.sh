#!/bin/bash
#
# trap EXIT
# trap INT
# trap TERM
# trap ERR
#

echo "pid is $$"

for i in {0..3}; do
	echo $i
	false
	sleep 1
done

trap "echo ..EXIT triggered.." EXIT
trap "echo ..INT triggered.."  INT
trap "echo ..TERM triggered.." TERM
trap "echo ..ERR triggered.."  ERR

for i in {0..3}; do
	echo $i
	false
	#if (( $i == 2 )); then exit 42; fi
	sleep 1
done

trap - ERR

for i in {0..3}; do
	echo $i
	false
	sleep 1
done

