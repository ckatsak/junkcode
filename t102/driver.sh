#!/usr/bin/env bash

PROMPT="[BashDriver]:"
PYTHON="/usr/bin/python"
FIFO1="./first.fifo"
FIFO2="./second.fifo"


echo "$PROMPT Spawning"
echo "$PROMPT Creating fifoz"
mkfifo $FIFO1
mkfifo $FIFO2
echo "$PROMPT Fifoz ready"

$PYTHON ./s.py &
S_PID=$!
sleep 2
if [[ $RANDOM < $((32768 / 2)) ]]; then # 0 <= $RANDOM < 32768
	echo "$PROMPT TERM < HOOK"
	kill -s TERM $S_PID
	sleep 0.3
	$PYTHON ./h.py &
	H_PID=$!
else
	echo "$PROMPT HOOK < TERM"
	$PYTHON ./h.py &
	H_PID=$!
	sleep 0.3
	kill -s TERM $S_PID
fi

trap "rm -vf $FIFO1 $FIFO2; kill -s KILL $S_PID; kill -s KILL $H_PID; exit 66;" EXIT
echo "$PROMPT Waiting..."
wait

echo "$PROMPT Cleaning up fifoz"
rm -vf $FIFO1 $FIFO2
echo "$PROMPT Fifoz removed"
echo "$PROMPT Exiting"
trap "" EXIT
