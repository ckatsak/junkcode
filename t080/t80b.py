#!/usr/bin/env python

import time
from pprint import pprint

import grequests


def print_res(res):
    pprint(res)


req = grequests.get('http://eu.httpbin.org/get', hooks={'response': print_res})
job = grequests.send(req, grequests.Pool(1))

for i in range(10):
    time.sleep(1)
    print i
