#!/bin/bash
#
# trap "" ERR	: comment in and out the `false` and observe

booh() {
	echo "IT'S A TRAP!"
}

echo "BEFORE"

trap "booh ; exit 42" ERR
true
#false
true
trap - ERR

true
false

echo "AFTER"
