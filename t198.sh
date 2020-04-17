#!/bin/bash
#
# Toy autoclicker
#
# ckatsak, Fri Apr 17 06:38:29 EEST 2020

set -o pipefail

print_logo() {
	echo
	echo "    __   _______ _  __    _     "
	echo "    \ \ / / ____| |/ /   / \    "
	echo "     \ V /|  _| | ' /   / _ \   "
	echo "      | | | |___| . \  / ___ \  "
	echo "      |_| |_____|_|\_\/_/   \_\ "
	echo
}

ts() {
	echo "[$(echo $(date) | cut -d' ' -f4)]"
}

log() {
	echo -e "$(ts) $1"
}

[[ $# -lt 1 || $# -gt 2 ]] && echo -e \
	"Usage:\n\t\$ $BASH_SOURCE <period_sec> [<init_sec>=10]" && exit 1
PERIOD=$1
INIT=10 && [[ ! "$2" == "" ]] && INIT=$2

XDOTOOL=$(which xdotool)
[[ -z $XDOTOOL ]] && echo -e "xdotool was not found in PATH." && exit 2

print_logo

log "Waiting $INIT sec to let you choose the mouse location & window to lock..."
sleep $INIT && eval $($XDOTOOL getmouselocation --shell)
log "Mouse location locked: ($X, $Y) @ window $WINDOW; clicking every $PERIOD sec...\n"

while sleep $PERIOD; do
	log "Left clicking ($X, $Y) @ window $WINDOW now!"
	$XDOTOOL mousemove --window $WINDOW $X $Y click 1 &
done
