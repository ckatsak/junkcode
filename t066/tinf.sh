#!/bin/bash
#
# tinf.sh

exec 1>&2

echo "[$$] TINF: start"
echo "[$$] TINF: my parent is $PPID"


LIMIT=10
i=0
while :; do
	sleep 1
	i=$(($i + 1))
	echo "[$$] TINF: $i"

	if (( $i == $LIMIT )); then
		break
	fi
done

echo "[$$] TINF: sending SIGUSR1 to my parent $PPID"
kill -SIGUSR1 $PPID
sleep 10


echo "[$$] TINF: end"
