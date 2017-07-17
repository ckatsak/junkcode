#!/usr/bin/env python

import os
from random import choice
from string import ascii_lowercase
import sys

if len(sys.argv) != 2:
    os.write(2, "Usage: %s <msg_size_2exp>\n" % sys.argv[0])

msg_size = 2 ** int(sys.argv[1])
bal = [c.encode('ascii') for c in ascii_lowercase]

with open('./msg.go', 'w') as gout:
    gout.writelines(["package sha256\n\nconst msg = \"",
                     b''.join([choice(bal) for _ in xrange(msg_size)]),
                     "\""])
