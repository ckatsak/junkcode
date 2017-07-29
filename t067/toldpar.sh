#!/bin/bash

echo "[$$]: Current jobs: $(jobs -p)"

{ sleep 15; echo "[$$]: END"; } &
echo "[$$]: Current jobs: $(jobs -p)"
sleep 2
echo

{ sleep 15; echo "[$$]: END"; } &
echo "[$$]: Current jobs: $(jobs -p)"
sleep 2
echo

{ sleep 15; echo "[$$]: END"; } &
echo "[$$]: Current jobs: $(jobs -p)"
sleep 2
echo

{ sleep 15; echo "[$$]: END"; } &
echo "[$$]: Current jobs: $(jobs -p)"
sleep 2
echo

{ sleep 15; echo "[$$]: END"; } &
echo "[$$]: Current jobs: $(jobs -p)"
sleep 2
echo

{ sleep 15; echo "[$$]: END"; } &
echo "[$$]: Current jobs: $(jobs -p)"
sleep 2
echo

echo "About to wait!"
wait
echo "Back from waiting!"
sleep 1
echo
echo "[$$]: Current jobs: $(jobs -p)"
