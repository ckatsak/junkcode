#!/usr/bin/env python
# -*- coding: utf-8 -*-

import gevent
import random
import time


SLEEP_FACTOR = 2


def task(tid):
    t = random.random() * SLEEP_FACTOR
    gevent.sleep(t)
    print "Task %s done after %f seconds" % (tid, t)


def sync():
    for i in xrange(10):
        task(i)


def async():
    greenlets = [gevent.spawn(task, i) for i in xrange(10)]
    gevent.joinall(greenlets)


print "Synchronous:"
start = time.time()
sync()
print "Synchronous total time: %f" % (time.time() - start)


print "Asynchronous:"
start = time.time()
async()
print "Asynchronous total time: %f" % (time.time() - start)
