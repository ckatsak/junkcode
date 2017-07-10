#!/usr/bin/env python


import errno
import os
import os.path
from pprint import pprint

WALKER = os.walk('.',
                 topdown=False,
                 onerror=lambda err: (_ for _ in ()).throw(err))


def hexdigits():
    hexdigits = ['a', 'b', 'c', 'd', 'e', 'f']
    hexdigits.extend(map(str, range(10)))
    for a in hexdigits:
        yield a


def mkdirs_rec(root, level):
    if level == 0:
        return
    for a in hexdigits():
        os.mkdir(os.path.join(root, a))
        mkdirs_rec(os.path.join(root, a), level-1)


def main():
    try:
        mkdirs_rec('.', 2)
    except OSError as err:
        if err.errno != errno.EEXIST:
            raise err
    # for (rootdir, dirnames, filenames) in WALKER.next():
    #     pprint((rootdir, dirnames, filenames))

    '''
    try:
        i = 0
        res = []
        while True:
            i += 1
            print '---' * 10, i, '---' * 10
            # pprint(WALKER.next())
            rootdir, dirnames, filenames = WALKER.next()
            # pprint((rootdir, dirnames, filenames))
            res.extend(filenames)
    except StopIteration as err:
        print err
    pprint(res)
    '''

    i = 0
    res = []
    for a, b in Walker('.'):
        pprint((a, b))
        res.extend(b)
        i += 1
        print i
    pprint(res)
    print '---' * 25

    w = iter(Walker('.'))
    pprint(w.next())
    pprint(w.next())
    pprint(w.next())
    pprint(w.next())
    print '---' * 25

    w = Walker('.')
    try:
        pprint(w())
        pprint(w())
        pprint(w())
        pprint(w())
        pprint(w())
        pprint(w())
        pprint(w())
        pprint(w())
        pprint(w())
        pprint(w())
        pprint(w())
        pprint(w())
    except StopIteration:
        print "Caught StopIteration!"


class Walker(object):
    """

    """
    def __init__(self, rootdir):
        self._rootdir = rootdir
        self._iter = iter(self)

    def __iter__(self):
        for rootdir, dirnames, filenames in \
                os.walk(self._rootdir, topdown=True,
                        onerror=lambda err: (_ for _ in ()).throw(err)):
            dirnames.sort()
            '''
            If we remove len(filenames) condition, and just sort filenames and
            yield, it is possible to guarantee for the total number of yielded
            results.
            In other words, using the condition len(filenames) we can optimize
            the performance of the walk, at the cost of more complex handling
            on the frontend. However, removing the condition, the steps of the
            walk are independent of the actual contents of the filesystem: each
            step walks a subdirectory, which means that, since all backend pods
            have the same subdirectory schema, they can sort of synchronize.
            '''
            if len(filenames) > 0:
                filenames.sort()
                yield rootdir, filenames

    def __call__(self):
        return self._iter.next()


if __name__ == '__main__':
    main()
