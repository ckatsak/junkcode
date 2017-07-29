#!/usr/bin/env python

import subprocess

for i in range(256):
    addr = "147.102.13." + str(i)
    try:
        output = subprocess.check_output(["host", addr])
        print addr, "-->", output.split()[-1]
    except subprocess.CalledProcessError:
        print '*' * 3, "No record for " + addr, '*' * 3
