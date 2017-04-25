#!/usr/bin/env python

import sys
import threading

import srvlookup


ind = 1


def lookup():
    res = srvlookup.lookup(sys.argv[1], sys.argv[2], sys.argv[3])

    global ind
    # r = res[1 + ind % 4]
    # ind += 1
    r = res[ind + 1]
    ind = (ind + 1) % 4
    print "[Thread %s]: %s:%d" % (
            threading.currentThread().name, r.host, r.port)


def loop_lookup():
    for i in range(10):
        lookup()


if __name__ == '__main__':
    if len(sys.argv) != 4:
        print "Usage:\n\t$ %s <svc> <proto> <domain>\n" % sys.argv[0]
        sys.exit(1)
    # res = srvlookup.lookup(sys.argv[1], sys.argv[2], sys.argv[3])
    #
    # from pprint import pprint
    # pprint(res)
    # for r in res:
    #     print "%s:%d" % (r.host, r.port)

    # loop_lookup()

    threads = [threading.Thread(
            target=loop_lookup, name=t) for t in range(1, 5)]
    for t in threads:
        t.start()
    for t in threads:
        t.join()
