# -*- coding: utf-8 -*-
"""

"""


import os.path


def inner_top_subdirs(prefix, first_obj, last_obj, index_height=3):
    start = int(''.join(first_obj[:index_height]), 16)
    end = int(''.join(last_obj[:index_height]), 16)

    for subdir_hex, useful in _inner_top_subdirs(start, end, index_height):
        yield os.path.join(prefix, *subdir_hex[:useful])


its = inner_top_subdirs


def _inner_top_subdirs(start, end, index_height):
    if start > end:
        end_aux = end + (1 << 4 * index_height)
    else:
        end_aux = end

    for i in _inner_top_subdirs_middle(
            start, end_aux, index_height-1, index_height):
        yield i


def _inner_top_subdirs_middle(start, end, exp, index_height):
    step = 1 << 4 * exp

    if end - start <= step:
        for i in _inner_top_subdirs_middle(start, end, exp-1, index_height):
            yield i
        raise StopIteration

    n = (start / step + 1) * step

    # left:
    for i in _inner_top_subdirs_left(start, n, exp-1, index_height):
        yield i

    # top-inner:
    while n <= end - step:
        yield hex(n)[2:].zfill(index_height+1)[1:], index_height-exp
        n += step
    last = n

    # right:
    for i in _inner_top_subdirs_right(last-1, end, exp-1, index_height):
        yield i


def _inner_top_subdirs_left(start, left_limit, exp, index_height):
    if exp < 0:
        raise StopIteration

    step = 1 << 4 * exp
    n = (start / step + 1) * step

    for i in _inner_top_subdirs_left(start, n, exp-1, index_height):
        yield i

    while n < left_limit:
        yield hex(n)[2:].zfill(index_height+1)[1:], index_height-exp
        n += step


def _inner_top_subdirs_right(right_limit, end, exp, index_height):
    if exp < 0:
        raise StopIteration

    step = 1 << 4 * exp
    n = (right_limit / step + 1) * step

    while n <= end - step:
        yield hex(n)[2:].zfill(index_height+1)[1:], index_height-exp
        n += step

    for i in _inner_top_subdirs_right(n-1, end, exp-1, index_height):
        yield i
