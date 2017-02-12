#!/bin/bash

exec 2>"./output.log"
exec 1>&2

echo "[$$]: Current jobs: $(jobs -p)"

./tch.sh &
echo "[$$]: Current jobs: $(jobs -p)"
sleep 2

./tch.sh &
echo "[$$]: Current jobs: $(jobs -p)"
sleep 2

./tch.sh &
echo "[$$]: Current jobs: $(jobs -p)"
sleep 2

./tch.sh &
echo "[$$]: Current jobs: $(jobs -p)"
sleep 2

./tch.sh &
echo "[$$]: Current jobs: $(jobs -p)"
sleep 2

./tch.sh &
echo "[$$]: Current jobs: $(jobs -p)"
sleep 2

echo "[$$]: About to wait!"
wait
echo "[$$]: Back from waiting!"
sleep 1
echo "[$$]: Current jobs: $(jobs -p)"
