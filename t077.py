#!/usr/bin/env python
"""
Use os.walk()'s onerror parameter, using a lambda which just raises whatever
OSError was actually raised and silented by os.walk(). Example demonstration:

    $ python t77.py /

Soon after execution a EACCES will be raised, as this script (if it isn't run
as root) doesn't have the permissions to access /.
"""

import os
import sys

all_files = []
for dirpath, subdirs, files in \
        os.walk(sys.argv[1], onerror=lambda err: (_ for _ in ()).throw(err)):
    all_files.extend(files)

print(all_files)
