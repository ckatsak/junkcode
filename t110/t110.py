#!/usr/bin/env python
# -*- coding: utf-8 -*-
"""
    gevent.pool example found at https://sdiehl.github.io/gevent-tutorial/

    Better run:
        $ strace -o pt110 -ff python t110.py
"""

from __future__ import print_function
from multiprocessing.pool import Pool as mPool
import time

from gevent.pool import Pool as gPool


def echo(i):
    time.sleep(0.001)
    return i


# Non Deterministic Process Pool

p = mPool(10)
run1 = [a for a in p.imap_unordered(echo, xrange(10))]
run2 = [a for a in p.imap_unordered(echo, xrange(10))]
run3 = [a for a in p.imap_unordered(echo, xrange(10))]
run4 = [a for a in p.imap_unordered(echo, xrange(10))]

print("mPool:", run1, run2, run3, run4)


# Deterministic Gevent Pool

p = gPool(10)
run1 = [a for a in p.imap_unordered(echo, xrange(10))]
run2 = [a for a in p.imap_unordered(echo, xrange(10))]
run3 = [a for a in p.imap_unordered(echo, xrange(10))]
run4 = [a for a in p.imap_unordered(echo, xrange(10))]

print("gPool:", run1, run2, run3, run4)
