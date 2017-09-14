#!/usr/bin/env python
# -*- coding: utf-8 -*-

import os.path

PREFIX = '/some/prefix'


hd = map(lambda x: str(hex(x))[2], range(16))
# hdd = {h: d for h, d in zip(hd, range(16))}


def sdr(prefix, first, last):
    """subdirs_range generator"""
    start = int(first[0]+first[1]+first[2], 16)
    end = int(last[0]+last[1]+last[2], 16)
    d2, d3 = divmod(end, 16)  # d2, d3 = end // 16, end % 16
    d1, d2 = divmod(d2, 16)   # d1, d2 = d2 // 16, d2 % 16
    for i in xrange(start, end+1):
        d2, d3 = divmod(i, 16)   # d2, d3 = i // 16, i % 16
        d1, d2 = divmod(d2, 16)  # d1, d2 = d2 // 16, d2 % 16
        yield os.path.join(prefix, str(hex(d1))[2],
                           str(hex(d2))[2], str(hex(d3))[2])


def sdrc(prefix, first, last):
    """cyclic subdirs_range generator"""
    start = int(first[0]+first[1]+first[2], 16)
    end = int(last[0]+last[1]+last[2], 16)
    if start <= end:
        for subdir in sdr(prefix, first, last):
            yield subdir
    else:
        for subdir in sdr(prefix, first, ('f', 'f', 'f')):
            yield subdir
        for subdir in sdr(prefix, ('0', '0', '0'), last):
            yield subdir
