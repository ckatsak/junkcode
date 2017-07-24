#!/usr/bin/env python

"""
    Quick comparison of append+sort vs bisect.insort.
"""

import timeit

LIST_SIZE = 2**16

print min(timeit.repeat('l=range(%d); l.append(844); l.sort()' % LIST_SIZE,
                        '',
                        repeat=7, number=1000))
print min(timeit.repeat('insort(range(%d), 844)' % LIST_SIZE,
                        'from bisect import insort',
                        repeat=7, number=1000))
