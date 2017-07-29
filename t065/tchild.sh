#!/bin/bash

echo "[$$]: CHILD start"
echo "[$$]: CHILD my parent is $PPID"

for i in {0..2}; do
	echo "[$$]: CHILD $i"
	sleep 1
done
echo
for i in {0..2}; do
	echo "[$$]: CHILD $i"
	sleep 1
done
echo

echo "[$$]: CHILD sending SIGUSR1 to $PPID"
kill -SIGUSR1 $PPID

echo "[$$]: CHILD end"
