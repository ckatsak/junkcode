#!/bin/bash
#
# tparent.sh

usr1_handler() {
	echo "[$$] PARENT: SIGUSR1 received"
}

exec 1>&2


echo "[$$] PARENT: start"


trap "usr1_handler" USR1
trap ':' USR1

./tinf.sh &
echo "[$$] PARENT: spawned TINF $!"

#./tchild.sh &
#echo "[$$] PARENT: spawned CHILD $!"

#while :; do
#	sleep 1
#done
#wait

#while :; do
#	echo "[$$] PARENT: duh!"
#	sleep 2 &
	wait $!
#done

trap - USR1


echo "[$$] PARENT: end"
