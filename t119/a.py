# -*- coding: utf-8 -*-
"""

"""

import os.path


def inner_top_subdirs(prefix, first_obj, last_obj, index_height=3):
    ''''''
    start = int(''.join(first_obj[:index_height]), 16)
    end = int(''.join(last_obj[:index_height]), 16)

    # max_obj_l = long('f' * len(first_obj), 16)
    # start_obj_l = long(first_obj, 16)
    # # start_plus_1_s = hex(start_obj_l+1)[2:-1].zfill(len(first_obj))
    # # start_plus_2_s = hex(start_obj_l+2)[2:-1].zfill(len(first_obj))
    # start_obj_plus_1_s = hex((start_obj_l+1) % max_obj_l)[2:-1].zfill(
    #         len(first_obj))
    # start_obj_plus_2_s = hex((start_obj_l+2) % max_obj_l)[2:-1].zfill(
    #         len(first_obj))
    if start == end or (start + 1) % int('f' * index_height, 16) == end:
        # Handle easy edge cases not handled by _tis().
        raise StopIteration
    elif (start + 2) % int('f' * index_height, 16) == end:
        # Handle easy edge case not handled by _tis().
        yield os.path.join(prefix, *hex(
            (start + 1) % int('f' * index_height, 16))[2:])
    elif start < end:
        for subdir_hex, useful in _tis(start, end, index_height):
            yield os.path.join(prefix, *subdir_hex[:useful])
    else:
        for subdir_hex, useful in _tis(
                start, int('f' * index_height, 16) + 1, index_height):
            yield os.path.join(prefix, *subdir_hex[:useful])
        yield os.path.join(prefix, *('0' * index_height))
        for subdir_hex, useful in _tis(0, end, index_height):
            yield os.path.join(prefix, *subdir_hex[:useful])


def _tis(start, end, index_height=3):
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
