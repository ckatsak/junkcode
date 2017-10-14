# -*- coding: utf-8 -*-
"""

"""

# import os.path


def tis(prefix, start, end, index_height=3):
    '''TODO'''
    # exp = 1 << 4 * (index_height - 1)
    max_exp = index_height - 1

    max_step = 1 << 4 * max_exp
    first = num = (start / max_step + 1) * max_step
    while num < end - max_step:
        print hex(num)[2:]
        num += max_step
    # last = num - max_step
    last = num

    try:
        for i in _tis_left(start, first, max_exp-1):
            # yield i
            print 'OUTPUT:', i
    except TypeError:
        pass
    print '\nLEFT SIDE DONE. LET\'S GO RIGHT NOW\n'
    try:
        # for i in _tis_right(end, last+max_step, max_exp-1):
        for i in _tis_right(end, last-1, max_exp-1):
            # yield i
            print 'OUTPUT:', i
    except TypeError:
        pass


def _tis_left(start, first, exp):
    ''''''
    print '\nEntering _tis_left(%s, %s, %d)' % (
            hex(start)[2:], hex(first)[2:], exp)
    if exp < 0:
        raise StopIteration

    step = 1 << 4 * exp
    print 'new step = %d' % step

    new_first = num = (start / step + 1) * step
    print 'new_first = %s' % hex(new_first)[2:]

    print 'Calling _tis_left(%s, %s, %d)...' % (
            hex(start)[2:], hex(new_first)[2:], exp-1)
    for i in _tis_left(start, new_first, exp-1):
        # print 'OUTPUT:', i
        yield i

    # while num < first - step:
    while num < first:
        # print 'OUTPUT:', hex(num)[2:]
        yield hex(num)[2:]
        num += step


def _tis_right(end, last, exp):
    ''''''
    print '\nEntering _tis_right(%s, %s, %d)' % (
            hex(end)[2:], hex(last)[2:], exp)
    if exp < 0:
        raise StopIteration

    step = 1 << 4 * exp
    print 'new step = %d' % step

    num = (last / step + 1) * step
    print 'first = num = %s' % hex(num)[2:]
    # while num < end:
    while num <= end - step:
        yield hex(num)[2:]
        # print 'OUTPUT:', hex(num)[2:]
        num += step
    new_last = num

    print 'Calling _tis_right(%s, %s, %d)...' % (
            hex(end)[2:], hex(new_last-1)[2:], exp-1)
    for i in _tis_right(end, new_last-1, exp-1):
        yield i
        # print 'OUTPUT:', i
