#!/bin/bash
#
# Redirect script's stdout and stderr to a file.
# Only for bash.

# Redirect stdout to $PWD/lala.txt.
exec >lala.txt

# Redirect stderr to stdout stream.
exec 2>&1

echo "O U T P U T -- stdOUT" >&1
echo "O U T P U T -- stdERR" >&2
