#!/usr/bin/env python

import grequests as gr


NUM_REQUESTS = 5
STOP_AT = 3


reqs = (
        gr.get(url) for url in (
            'http://httpbin.org/get?i=%d' % (i)
            for i in xrange(NUM_REQUESTS)
        )
)

print "Sending %d requests... Gonna wait for the first %d responses..." % (
        NUM_REQUESTS, STOP_AT)

n = 0
for r in gr.imap(reqs, size=NUM_REQUESTS):
    print r
    if r.status_code == 200:
        n += 1
    if n == STOP_AT:
        break
