#!/usr/bin/env python
# -*- coding: utf-8 -*-
"""

"""

from __future__ import print_function
import json
import os
import signal
import sys
import threading as thr
import time


log = None

FIFO1 = 'first.fifo'
FIFO2 = 'second.fifo'


def term_handler(signum, frame):
    print("[%s]: Wild SIGTERM appears!" % thr.current_thread().getName(),
          file=log)
    print("[%s]: Opening '%s' for writing..." % (
          thr.current_thread().getName(), FIFO1), file=log)
    with open(FIFO1, 'w') as fifo1:
        print("[%s]: Just opened '%s' for writing" % (
              thr.current_thread().getName(), FIFO1), file=log)
        snd = json.dumps({
            'start': 'ABC',
            'end': 'XYZ',
            'dest': 'CHRISTOULAS',
        })
        print("[%s]: Sending '%s'" % (thr.current_thread().getName(), snd),
              file=log)
        fifo1.write(snd)
        print("[%s]: Closing '%s' (writing)" % (thr.current_thread().getName(),
              FIFO1), file=log)


def exit_fn():
    print("[%s]: Spawning" % thr.current_thread().getName(), file=log)

    print("[%s]: Opening '%s' for reading..." % (
          thr.current_thread().getName(), FIFO2), file=log)
    with open(FIFO2, 'r'):
        print("[%s]: Just opened '%s' for reading" % (
              thr.current_thread().getName(), FIFO2), file=log)
        print("[%s]: Closing '%s' (reading)" % (thr.current_thread().getName(),
              FIFO2), file=log)

    print("[%s]: Shutting everything down..." % thr.current_thread().getName(),
          file=log)
    for i in xrange(4, 0, -1):
        time.sleep(0.4)
        print("[%s]: %d..." % (thr.current_thread().getName(), i), file=log)
        time.sleep(0.4)
    os._exit(42)


def main():
    print("[%s]: Spawning" % (thr.current_thread().getName()), file=log)
    signal.signal(signal.SIGTERM, term_handler)
    while True:
        pass


if __name__ == '__main__':
    log = sys.stderr
    thr.Thread(target=exit_fn, name='ExitThread').start()
    main()
