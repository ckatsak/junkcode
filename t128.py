#!/usr/bin/env python
"""
    Testing grequests.imap() using URLs at http://httpbin.org.
    Execution & output example:

        $ ./t128.py 11 1 delay
        INFO:RequestSender:Sending 11 requests and waiting for the first 1 responses.
        DEBUG:urllib3.connectionpool:Starting new HTTP connection (1): httpbin.org
        DEBUG:urllib3.connectionpool:Starting new HTTP connection (1): httpbin.org
        DEBUG:urllib3.connectionpool:Starting new HTTP connection (1): httpbin.org
        DEBUG:urllib3.connectionpool:Starting new HTTP connection (1): httpbin.org
        DEBUG:urllib3.connectionpool:Starting new HTTP connection (1): httpbin.org
        DEBUG:urllib3.connectionpool:Starting new HTTP connection (1): httpbin.org
        DEBUG:urllib3.connectionpool:Starting new HTTP connection (1): httpbin.org
        DEBUG:urllib3.connectionpool:Starting new HTTP connection (1): httpbin.org
        DEBUG:urllib3.connectionpool:Starting new HTTP connection (1): httpbin.org
        DEBUG:urllib3.connectionpool:Starting new HTTP connection (1): httpbin.org
        DEBUG:urllib3.connectionpool:Starting new HTTP connection (1): httpbin.org
        DEBUG:urllib3.connectionpool:http://httpbin.org:80 "GET /delay/0 HTTP/1.1" 200 316
        INFO:RequestSender:<Response [200]> after 0.528023958206 seconds
        DEBUG:urllib3.connectionpool:http://httpbin.org:80 "GET /delay/1 HTTP/1.1" 200 316
        DEBUG:urllib3.connectionpool:http://httpbin.org:80 "GET /delay/2 HTTP/1.1" 200 316
        DEBUG:urllib3.connectionpool:http://httpbin.org:80 "GET /delay/3 HTTP/1.1" 200 316
        DEBUG:urllib3.connectionpool:http://httpbin.org:80 "GET /delay/4 HTTP/1.1" 200 316
        DEBUG:urllib3.connectionpool:http://httpbin.org:80 "GET /delay/5 HTTP/1.1" 200 316
        DEBUG:urllib3.connectionpool:http://httpbin.org:80 "GET /delay/6 HTTP/1.1" 200 316
        Result: True

        $ ./t128.py 11 9 delay
        INFO:RequestSender:Sending 11 requests and waiting for the first 9 responses.
        DEBUG:urllib3.connectionpool:Starting new HTTP connection (1): httpbin.org
        DEBUG:urllib3.connectionpool:Starting new HTTP connection (1): httpbin.org
        DEBUG:urllib3.connectionpool:Starting new HTTP connection (1): httpbin.org
        DEBUG:urllib3.connectionpool:Starting new HTTP connection (1): httpbin.org
        DEBUG:urllib3.connectionpool:Starting new HTTP connection (1): httpbin.org
        DEBUG:urllib3.connectionpool:Starting new HTTP connection (1): httpbin.org
        DEBUG:urllib3.connectionpool:Starting new HTTP connection (1): httpbin.org
        DEBUG:urllib3.connectionpool:Starting new HTTP connection (1): httpbin.org
        DEBUG:urllib3.connectionpool:Starting new HTTP connection (1): httpbin.org
        DEBUG:urllib3.connectionpool:Starting new HTTP connection (1): httpbin.org
        DEBUG:urllib3.connectionpool:Starting new HTTP connection (1): httpbin.org
        DEBUG:urllib3.connectionpool:http://httpbin.org:80 "GET /delay/0 HTTP/1.1" 200 316
        INFO:RequestSender:<Response [200]> after 0.69969201088 seconds
        DEBUG:urllib3.connectionpool:http://httpbin.org:80 "GET /delay/1 HTTP/1.1" 200 316
        DEBUG:urllib3.connectionpool:http://httpbin.org:80 "GET /delay/2 HTTP/1.1" 200 316
        DEBUG:urllib3.connectionpool:http://httpbin.org:80 "GET /delay/3 HTTP/1.1" 200 316
        DEBUG:urllib3.connectionpool:http://httpbin.org:80 "GET /delay/4 HTTP/1.1" 200 316
        DEBUG:urllib3.connectionpool:http://httpbin.org:80 "GET /delay/5 HTTP/1.1" 200 316
        DEBUG:urllib3.connectionpool:http://httpbin.org:80 "GET /delay/6 HTTP/1.1" 200 316
        INFO:RequestSender:<Response [200]> after 12.7020070553 seconds
        INFO:RequestSender:<Response [200]> after 24.7029941082 seconds
        INFO:RequestSender:<Response [200]> after 36.7135109901 seconds
        INFO:RequestSender:<Response [200]> after 48.7244970798 seconds
        INFO:RequestSender:<Response [200]> after 60.7355630398 seconds
        INFO:RequestSender:<Response [200]> after 72.7375080585 seconds
        CRITICAL:RequestSender:RequestSender.exception_handler(): <grequests.AsyncRequest object at 0x7f959bb525d0> raised HTTPConnectionPool(host='httpbin.org', port=80): Read timed out. (read timeout=7)
        rTimeout caught after 84.7488391399 seconds
        Result: False
"""

import logging
import sys
import time

import grequests
from requests import Timeout as rTimeout


CONN_TIMEOUT_SEC = 3.05
READ_TIMEOUT_SEC = 7


class RequestSender(object):
    def __init__(self, num_reqs, stop_at):
        self.num_reqs = num_reqs
        self.stop_at = stop_at
        self.logger = logging.getLogger('RequestSender')

    def exception_handler(self, request, exception):
        self.logger.critical("RequestSender.exception_handler(): %s raised %s",
                             request, exception)
        raise exception

    def send_requests(self, requests):
        n = 0
        self.logger.info(
                "Sending %d requests and waiting for the first %d responses.",
                self.num_reqs, self.stop_at)
        start = time.time()
        try:
            for r in grequests.imap(requests,
                                    size=self.num_reqs,
                                    exception_handler=self.exception_handler):
                self.logger.info("%s after %s seconds", r, time.time() - start)
                if r.status_code == 200:
                    time.sleep(12)
                    n += 1
                if n == self.stop_at:
                    return True
        except rTimeout:
            self.logger.warning("rTimeout caught after %s seconds",
                                time.time() - start)
            return False
        self.logger.info("...Done after %s seconds", time.time() - start)
        return True


def main():
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

    requests = (
            grequests.get(url, timeout=(CONN_TIMEOUT_SEC, READ_TIMEOUT_SEC))
            for url in (
                TARGET_URLS[OPTION.lower()] % (i)
                for i in xrange(NUM_REQUESTS)
            )
    )

    res = RequestSender(NUM_REQUESTS, STOP_AT).send_requests(requests)
    print "Result:", res


if __name__ == '__main__':
    logging.basicConfig(level=logging.DEBUG)
    main()
