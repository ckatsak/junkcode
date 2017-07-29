#!/usr/bin/env python

import os.path
import random
import time


class ListSession(object):
    def __init__(self, handle, shard_gen, subdir_gen):
        self.handle = handle
        self._shard_gen = shard_gen
        self._subdir_gen = subdir_gen

    def __iter__(self):
        return self

    def __next__(self):
        for shard in self._shard_gen:
            for subdir in self._subdir_gen:
                return shard, subdir
        raise StopIteration

    next = __next__


def next_subdir():
    hd = map(lambda x: str(hex(x))[2], range(16))
    for d1 in hd:
        for d2 in hd:
            yield os.path.join(d1, d2)


def next_shard():
    for i in range(5):
        yield i


def main():
    ls = ListSession(42, (i for i in xrange(20)), next_subdir())
    for a, b in ls:
        print 'a:', a, ', b:', b
    # while True:
    #     try:
    #         print next(ls)
    #     except StopIteration:
    #         print "StopIteration!"
    #         break
    print "THE END"


class LS2(object):
    def __init__(self, handle, max_shard, subdir_gen_func):
        self.handle = handle
        self.max_shard = max_shard
        self._subdir_gen_func = subdir_gen_func
        self._subdir_gen = self._subdir_gen_func()

    def __iter__(self):
        return self

    def __next__(self):
        for shard in xrange(self.max_shard):
            try:
                for subdir in self._subdir_gen:
                    return shard, subdir
            except StopIteration:
                print "__next__(): Caught StopIteration!"
                if shard != self.max_shard:
                    self._subdir_gen = self._subdir_gen_func()
                    continue

    next = __next__


def main2():
    for a, b in LS2(42, 3, next_subdir):
        print 'a:', a, ', b:', b
    print "THE END"


class LS3(object):
    def __init__(self, handle, subdir_gen, max_shard):
        self.handle = handle
        self._subdir_gen = subdir_gen
        self._curr_subdir = None
        self.max_shard = max_shard
        self._curr_shard = 0

    def __iter__(self):
        return self

    def __next__(self):
        self._curr_shard += 1
        if self._curr_shard == self.max_shard or self._curr_subdir is None:
            self._curr_shard = 0
            self._curr_subdir = next(self._subdir_gen)
        return self._curr_subdir, self._curr_shard
    next = __next__


def main3():
    # for a, b in LS3(42, next_subdir(), 2):

    # for a, b in LS3(42, next_subdir(), 3):
    #     print "subdir:", a, ", shard:", b
    # print "THE END"
    aux31(LS3(42, next_subdir(), 3), 10)


def aux31(ls, times):
    for i in xrange(times):
        print next(ls)
        time.sleep(random.random() * 7)


if __name__ == '__main__':
    # main()
    # main2()
    main3()
