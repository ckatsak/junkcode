#!/usr/bin/env python
"""
    Testing grequests.imap() using URLs at http://httpbin.org.
    Execution & output example:

        $ ./e.py 11 1 delay
        Sending 11 requests... Gonna wait for the first 1 responses...
        <Response [200]> after 0.508654117584 seconds
        Result: True

        $ ./e.py 11 9 delay
        Sending 11 requests... Gonna wait for the first 9 responses...
        <Response [200]> after 0.526062965393 seconds
        <Response [200]> after 12.5303828716 seconds
        <Response [200]> after 24.5402090549 seconds
        <Response [200]> after 36.5510919094 seconds
        <Response [200]> after 48.5614550114 seconds
        <Response [200]> after 60.5719809532 seconds
        <Response [200]> after 72.5738158226 seconds
        rTimeout caught after 84.5823879242 seconds
        Result: False
"""

import sys
import time

import grequests
from requests import Timeout as rTimeout


CONN_TIMEOUT_SEC = 3.05
READ_TIMEOUT_SEC = 7


def exception_handler(request, exception):
    raise exception


def send_requests(reqs):
    global NUM_REQUESTS, STOP_AT
    n = 0
    print "Sending %d requests... Gonna wait for the first %d responses..." % (
            NUM_REQUESTS, STOP_AT)
    start = time.time()
    try:
        for r in grequests.imap(
                reqs, size=NUM_REQUESTS, exception_handler=exception_handler):
            print "%s after %s seconds" % (r, time.time() - start)
            if r.status_code == 200:
                time.sleep(12)
                n += 1
            if n == STOP_AT:
                return True
    except rTimeout:
        print >>sys.stderr, "rTimeout caught after %s seconds" % (
                time.time() - start)
        return False
    print "...Done after %s seconds" % (time.time() - start)
    return True


def main():
    if len(sys.argv) != 4:
        print >>sys.stderr, """
        \tUsage:\n\t\t$ %s <NUM_REQUESTS> <STOP_AT> <OPTION>\n
        \tOPTIONS:\n\t\tdelay\n\t\tparam\n""" % sys.argv[0]
        sys.exit(1)

    global NUM_REQUESTS, STOP_AT, OPTION
    NUM_REQUESTS = int(sys.argv[1])
    STOP_AT = int(sys.argv[2])
    OPTION = sys.argv[3].strip()

    TARGET_URLS = {
            'delay': 'http://httpbin.org/delay/%d',
            'param': 'http://httpbin.org/get?param=%d',
    }

    requests = (
            grequests.get(url, timeout=(CONN_TIMEOUT_SEC, READ_TIMEOUT_SEC))
            for url in (
                TARGET_URLS[OPTION.lower()] % (i)
                for i in xrange(NUM_REQUESTS)
            )
    )

    res = send_requests(requests)
    print "Result:", res


if __name__ == '__main__':
    main()
