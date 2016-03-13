#!/usr/bin/env python3
#
# script -c "./t48a.py" /dev/null | ./t48b.py
#

import sys

def input_process():
    while True:
        print(sys.argv[0] + ": I just got " + input())

if __name__ == "__main__":
    input_process()
