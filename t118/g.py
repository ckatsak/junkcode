# -*- coding: utf-8 -*-
"""

"""

# import os.path


def tis(prefix, start, end, index_height=3):
    '''
    CORRECT; SORTED
    '''
    max_exp = index_height - 1
    max_step = 1 << 4 * max_exp
    first = num = (start / max_step + 1) * max_step

    # left:
    for i in _tis_left(start, first, max_exp-1, index_height):
        yield i

    # top-inner:
    while num < end - max_step:
        yield hex(num)[2:].zfill(index_height), index_height - max_exp
        num += max_step
    last = num

    # right:
    for i in _tis_right(end, last-1, max_exp-1, index_height):
        yield i


def _tis_left(start, first, exp, index_height):
    ''''''
    if exp < 0:
        raise StopIteration
    step = 1 << 4 * exp
    new_first = num = (start / step + 1) * step
    for i in _tis_left(start, new_first, exp-1, index_height):
        yield i
    while num < first:
        yield hex(num)[2:].zfill(index_height), index_height - exp
        num += step


def _tis_right(end, last, exp, index_height):
    ''''''
    if exp < 0:
        raise StopIteration
    step = 1 << 4 * exp
    num = (last / step + 1) * step

    while num <= end - step:
        yield hex(num)[2:].zfill(index_height), index_height - exp
        num += step
    new_last = num
    for i in _tis_right(end, new_last-1, exp-1, index_height):
        yield i
