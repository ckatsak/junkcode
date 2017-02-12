#!/bin/bash

usr1_handler() {
	echo "[$$]: PARENT SIGUSR1 received"
}


echo "[$$]: PARENT start"


trap "usr1_handler" USR1

./tchild.sh &
echo "[$$]: PARENT spawned $!"
#while :; do
#	sleep 1
#done
wait

trap - USR1


echo "[$$]: PARENT end"
