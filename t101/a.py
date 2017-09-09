#!/usr/bin/env python
# -*- coding: utf-8 -*-
"""
    Test blocking open(2) for named pipes.
"""

import os
import os.path
import signal
import threading
import time

NAMED_PIPE = 'yoyo.fifo'


def extra_thread_fn():
    print "[%s]: Spawning" % threading.current_thread().getName()

    # Stall 10s before opening the fifo
    for i in xrange(10):
        time.sleep(0.5)
        print "[%s]: %d..." % (threading.current_thread().getName(), i+1)
        time.sleep(0.5)

    print "[%s]: Opening '%s' for writing..." % (
            threading.current_thread().getName(), NAMED_PIPE)
    with open(NAMED_PIPE, 'w'):
        # Stall 10s before closing the fifo
        print "[%s]: Just opened '%s' for writing!" % (
                threading.current_thread().getName(), NAMED_PIPE)
        for i in xrange(10, 0, -1):
            time.sleep(0.5)
            print "[%s]: %d..." % (threading.current_thread().getName(), i)
            time.sleep(0.5)
        print "[%s]: Closing '%s'" % (
                threading.current_thread().getName(), NAMED_PIPE)


def term_handler(signum, frame):
    print "[%s]: A wild SIGTERM appears!" % (
            threading.current_thread().getName())


def main():
    threading.Thread(target=extra_thread_fn, name='extra_thread').start()

    print "[%s]: Installing signal handler for SIGTERM" % (
            threading.current_thread().getName())
    signal.signal(signal.SIGTERM, term_handler)

    if os.path.exists(NAMED_PIPE):
        print "[%s]: Named pipe '%s/%s' already exists" % (
                threading.current_thread().getName(), os.getcwd(), NAMED_PIPE)
    else:
        print "[%s]: Creating named pipe '%s/%s'" % (
                threading.current_thread().getName(), os.getcwd(), NAMED_PIPE)
        os.mkfifo(NAMED_PIPE)

    print "[%s]: Opening '%s' for reading..." % (
            threading.current_thread().getName(), NAMED_PIPE)
    with open(NAMED_PIPE, 'r'):
        print "[%s]: Just opened '%s' for reading!" % (
                threading.current_thread().getName(), NAMED_PIPE)
        print "[%s]: Closing '%s'" % (
                threading.current_thread().getName(), NAMED_PIPE)


if __name__ == '__main__':
    main()
