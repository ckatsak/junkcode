#!/bin/bash
#
# &, jobs, wait

QEMU=$(which which)
echo $?
echo $QEMU

echo "sleeping 1.."
sleep 7 &
echo "\$$ == $$"
echo -e "Jobs:\n$(jobs -p)"

echo "sleeping 2.."
sleep 5 &
echo "\$$ == $$"
echo -e "Jobs:\n$(jobs -p)"

wait
echo "quitting..."
