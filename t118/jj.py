# -*- coding: utf-8 -*-
"""

"""

# import os.path


def tis(prefix, start, end, index_height=3):
    '''
    NOT CORRECT; SORTED
    '''
    max_exp = index_height - 1
    max_step = 1 << 4 * max_exp
    num = (start / max_step + 1) * max_step
    # if num > max_step and max_step > end - start:
    if end - start > max_step:
        first = num
    else:
        first = num = start

    # left:
    for i in _tis_left(start, end, first, max_exp-1, index_height):
        yield i

    # top-inner:
    # if start > max_step:
    while num < end - max_step:
        yield hex(num)[2:].zfill(index_height), index_height - max_exp
        num += max_step
    last = num

    # right:
    for i in _tis_right(start, end, last-1, max_exp-1, index_height):
        yield i


def _tis_left(start, end, first, exp, index_height):
    ''''''
    tab = '\t' * (index_height - exp - 1)
    print tab, "Entering _tis_left(%s, %s, %d, %d)" % (
            hex(start)[2:], hex(first)[2:], exp, index_height)
    if exp < 0:
        print tab, "raising StopIteration"
        raise StopIteration
    step = 1 << 4 * exp
    new_first = num = (start / step + 1) * step
    # if num > step:
    # if end - start > step:
    if first - start > step:
        new_first = num
    else:
        new_first = num = start
    print tab, "step =", step
    print tab, "num =", hex(new_first)[2:]
    print tab, "up to", hex(first)[2:]

    for i in _tis_left(start, end, new_first, exp-1, index_height):
        yield i

    # if start > step:
    while num < first:
        yield hex(num)[2:].zfill(index_height), index_height - exp
        num += step
    print tab, "Leaving _tis_left(..., ..., %d, %d)" % (exp, index_height)


def _tis_right(start, end, last, exp, index_height):
    ''''''
    tab = '\t' * (index_height - exp - 1)
    print tab, "Entering _tis_right(%s, %s, %d, %d)" % (
            hex(end)[2:], hex(last)[2:], exp, index_height)
    if exp < 0:
        print tab, "raising StopIteration"
        raise StopIteration
    step = 1 << 4 * exp
    num = (last / step + 1) * step
    if num <= step:
        num = last
    print tab, "step =", step
    print tab, "num =", hex(num)[2:]
    print tab, "up to", hex(end - step)[2:]

    # if last > step:
    while num <= end - step:
        yield hex(num)[2:].zfill(index_height), index_height - exp
        num += step
    new_last = num
    for i in _tis_right(start, end, new_last-1, exp-1, index_height):
        yield i
    print tab, "Leaving _tis_right(..., ..., %d, %d)" % (exp, index_height)
