#!/usr/bin/env python
"""
    Testing grequests.imap() using URLs at http://httpbin.org.
    Execution & output example:

        $ ./c.py 11 9 delay
        Sending 11 requests... Gonna wait for the first 9 responses...
        <Response [200]> after 0.452767133713 seconds
        <Response [200]> after 1.59490013123 seconds
        <Response [200]> after 2.72903418541 seconds
        <Response [200]> after 3.9577050209 seconds
        <Response [200]> after 4.56124520302 seconds
        rTimeout caught after 7.30390810966 seconds
        ...Done after 7.30392813683 seconds
"""

import sys
import time

import grequests
from requests import Timeout as rTimeout


CONN_TIMEOUT_SEC = 3.05
READ_TIMEOUT_SEC = 7


def exception_handler(request, exception):
    raise exception


if len(sys.argv) != 4:
    print >>sys.stderr, """
    \tUsage:\n\t\t$ %s <NUM_REQUESTS> <STOP_AT> <OPTION>\n
    \tOPTIONS:\n\t\tdelay\n\t\tparam\n""" % sys.argv[0]
    sys.exit(1)

NUM_REQUESTS = int(sys.argv[1])
STOP_AT = int(sys.argv[2])
OPTION = sys.argv[3].strip()

TARGET_URLS = {
        'delay': 'http://httpbin.org/delay/%d',
        'param': 'http://httpbin.org/get?param=%d',
}

reqs = (
        grequests.get(url, timeout=(CONN_TIMEOUT_SEC, READ_TIMEOUT_SEC))
        for url in (
            TARGET_URLS[OPTION] % (i)
            for i in xrange(NUM_REQUESTS)
        )
)
n = 0
print "Sending %d requests... Gonna wait for the first %d responses..." % (
        NUM_REQUESTS, STOP_AT)
start = time.time()
try:
    for r in grequests.imap(
            reqs, size=NUM_REQUESTS, exception_handler=exception_handler):
        print "%s after %s seconds" % (r, time.time() - start)
        if r.status_code == 200:
            n += 1
        if n == STOP_AT:
            break
except rTimeout as timeout:
    print >>sys.stderr, "rTimeout caught after %s seconds" % (
            time.time() - start)
print "...Done after %s seconds" % (time.time() - start)
