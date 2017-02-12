#!/bin/bash
#
# bash arrays

WOO=3
TAPIFS=()

for i in $(seq 0 $(($WOO - 1))); do
	echo "i==$i"
	TAPIFS+=("tap$i ")
	echo "TAPIFS[$i]==${TAPIFS[$i]}"
	echo
done

echo TAPIFS=$TAPIFS

echo TAPIFS LOOP:
for tapif in "${TAPIFS[@]}"; do
	echo "$tapif"
done
