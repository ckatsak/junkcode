#!/usr/bin/env python3
#
# script -c "./t48a.py" /dev/null | ./t48b.py
#

import time
import random

def output():
    for i in range(15):
        print(random.randint(1, 101))
        time.sleep(3)

if __name__ == "__main__":
    output()
