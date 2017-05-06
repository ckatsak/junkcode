#!/usr/bin/env python

import grequests
from pprint import pprint


def print_res(res):
    pprint(res)


req = grequests.get('http://eu.httpbin.org/get',
                    hooks=dict(response=print_res))
res = grequests.map([req])
# pprint(res)

for i in range(10):
    print i
