#!/usr/bin/env python
"""
    Testing grequests.imap() using URLs at http://httpbin.org.
    Execution & output example:

        $ ./b.py 11 9 delay
        Sending 11 requests... Gonna wait for the first 9 responses...
        <Response [200]> after 0.489258050919 seconds
        <Response [200]> after 1.50731992722 seconds
        <Response [200]> after 2.52941703796 seconds
        <Response [200]> after 3.55319499969 seconds
        <Response [200]> after 4.57703089714 seconds
        <Response [200]> after 5.61269497871 seconds
        <Response [200]> after 6.52286291122 seconds
        Request to 'http://httpbin.org/delay/8' timed-out: HTTPConnectionPool(host='httpbin.org', port=80): Read timed out. (read timeout=7)
        Request to 'http://httpbin.org/delay/7' timed-out: HTTPConnectionPool(host='httpbin.org', port=80): Read timed out. (read timeout=7)
        Request to 'http://httpbin.org/delay/10' timed-out: HTTPConnectionPool(host='httpbin.org', port=80): Read timed out. (read timeout=7)
        Request to 'http://httpbin.org/delay/9' timed-out: HTTPConnectionPool(host='httpbin.org', port=80): Read timed out. (read timeout=7)
        ...Done after 8.29783391953 seconds
"""

import sys
import time

import grequests
from requests import Timeout as rTimeout


CONN_TIMEOUT_SEC = 3.05
READ_TIMEOUT_SEC = 7


def timeout_handler(request, exception):
    if isinstance(exception, rTimeout):
        print "'%s' request to '%s' timed-out: %s" % (
                request.method, request.url, exception)
    else:
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
        # grequests.get(url)
        for url in (
            TARGET_URLS[OPTION] % (i)
            for i in xrange(NUM_REQUESTS)
        )
)
n = 0
print "Sending %d requests... Gonna wait for the first %d responses..." % (
        NUM_REQUESTS, STOP_AT)
start = time.time()
# for r in grequests.imap(reqs, size=NUM_REQUESTS):
for r in grequests.imap(
        reqs, size=NUM_REQUESTS, exception_handler=timeout_handler):
    print "%s after %s seconds" % (r, time.time() - start)
    if r.status_code == 200:
        n += 1
    if n == STOP_AT:
        break
print "...Done after %s seconds" % (time.time() - start)
