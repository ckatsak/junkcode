#!/bin/bash
#
# pick random unused port

echo "User: $USER"

NETSTAT=$(which "netstat")
AWK=$(which "awk")

while [ -z $SSH_FW_HOSTPORT ]; do
	# Pick a random port in range [50000, 65536)
	SSH_FW_HOSTPORT=$((50000 + $(($RANDOM % 15536))))
	#SSH_FW_HOSTPORT=51177
	echo $SSH_FW_HOSTPORT
	# Get the currently in use tcp4 ports - might change between iterations
	PORTS_IN_USE=$("$NETSTAT" -an | "$AWK" '/^tcp /{ print $4 }' | "$AWK" -F ":" '{ print $2 }')
	for port in $PORTS_IN_USE; do
		# If our so far-chosen port is currently in use, repeat
		if [ "$port" == "$SSH_FW_HOSTPORT" ]; then
			echo "\$port == $port == \$SSH_FW_HOSTPORT"
			SSH_FW_HOSTPORT=""
			continue 2
		fi  
	done
done

