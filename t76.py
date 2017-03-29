#!/usr/bin/env python

import argparse
import sys

parser = argparse.ArgumentParser(description="Yo lelehson")

parser.add_argument('-a', action='store_true', dest='a', default=False)
parser.add_argument('--a', action='store_true', dest='a', default=False)
parser.add_argument('-b', action='store', dest='b')
parser.add_argument('-c', action='store', dest='c', type=int)

# print parser.parse_args(['-a', '-bval', '-c', '3'])
args = parser.parse_args(sys.argv[1:])
print "a =", args.a
print "b =", args.b
print "c =", args.c
