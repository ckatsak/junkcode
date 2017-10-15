# -*- coding: utf-8 -*-
"""

"""


def tis(prefix, start, end, index_height=3):
    '''
    LOOKS ACTUALLY CORRECT; SORTED
    '''
    max_exp = index_height - 1
    for i in _tis_middle(start, end, max_exp, index_height):
        yield i


def _tis_middle(start, end, exp, index_height):
    step = 1 << 4 * exp
    num = (start / step + 1) * step
    if end - start > step:
        first = num
    else:
        for i in _tis_middle(start, end, exp-1, index_height):
            yield i
        raise StopIteration

    # left:
    for i in _tis_left(start, end, first, exp-1, index_height):
        yield i

    # top-inner:
    while num < end - step:
        yield hex(num)[2:].zfill(index_height), index_height - exp
        num += step
    last = num

    # right:
    for i in _tis_right(start, end, last-1, exp-1, index_height):
        yield i


def _tis_left(start, end, first, exp, index_height):
    ''''''
    if exp < 0:
        raise StopIteration
    step = 1 << 4 * exp
    new_first = num = (start / step + 1) * step
    if first - start > step:
        new_first = num
    else:
        new_first = num = start

    for i in _tis_left(start, end, new_first, exp-1, index_height):
        yield i

    while num < first:
        yield hex(num)[2:].zfill(index_height), index_height - exp
        num += step


def _tis_right(start, end, last, exp, index_height):
    ''''''
    if exp < 0:
        raise StopIteration
    step = 1 << 4 * exp
    num = (last / step + 1) * step
    if num <= step:
        num = last

    while num <= end - step:
        yield hex(num)[2:].zfill(index_height), index_height - exp
        num += step
    new_last = num
    for i in _tis_right(start, end, new_last-1, exp-1, index_height):
        yield i
