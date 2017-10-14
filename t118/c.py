# -*- coding: utf-8 -*-
"""

"""

# import os.path


def tis(prefix, start, end, index_height=3):
    '''
    CORRECT; UNSORTED
    '''
    max_exp = index_height - 1

    max_step = 1 << 4 * max_exp
    first = num = (start / max_step + 1) * max_step
    while num < end - max_step:
        # print hex(num)[2:]
        yield hex(num)[2:]
        num += max_step
    last = num

    for i in _tis_left(start, first, max_exp-1):
        yield i
    for i in _tis_right(end, last-1, max_exp-1):
        yield i


def _tis_left(start, first, exp):
    ''''''
    if exp < 0:
        raise StopIteration
    step = 1 << 4 * exp
    new_first = num = (start / step + 1) * step
    for i in _tis_left(start, new_first, exp-1):
        yield i
    while num < first:
        yield hex(num)[2:]
        num += step


def _tis_right(end, last, exp):
    ''''''
    if exp < 0:
        raise StopIteration
    step = 1 << 4 * exp
    first = num = (last / step + 1) * step

    while num <= end - step:
        num += step
    new_last = num
    for i in _tis_right(end, new_last-1, exp-1):
        yield i

    num = first
    while num <= end - step:
        yield hex(num)[2:]
        num += step
