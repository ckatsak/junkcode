#!/bin/bash

exec 1>&2

for i in {1..10}; do
	sleep 1
	echo "[$$]: $i"
done
