#!/usr/bin/env python
# -*- coding: utf-8 -*-
"""

"""

from __future__ import print_function
import sys
import time


NAME = 'HookProcess'
log = None

FIFO1 = 'first.fifo'
FIFO2 = 'second.fifo'


def hook():
    print("[%s]: Spawning" % NAME, file=log)

    print("[%s]: Opening '%s' for reading..." % (NAME, FIFO1), file=log)
    with open(FIFO1, 'r') as fifo1:
        print("[%s]: Just opened '%s' for reading" % (NAME, FIFO1), file=log)
        inc = fifo1.read()
        print("[%s]: Received: '%s'" % (NAME, inc), file=log)
        for i in xrange(15, 0, -1):
            time.sleep(0.4)
            print("[%s]: %d..." % (NAME, i), file=log)
            time.sleep(0.4)
        print("[%s]: Closing '%s' (reading)" % (NAME, FIFO1), file=log)

    print("[%s]: Opening '%s' for writing..." % (NAME, FIFO2), file=log)
    with open(FIFO2, 'w'):
        print("[%s]: Just opened '%s' for writing" % (NAME, FIFO2), file=log)
        print("[%s]: Closing '%s' (writing)" % (NAME, FIFO2), file=log)

    print("[%s]: Exiting" % NAME, file=log)


if __name__ == '__main__':
    log = sys.stderr
    hook()
