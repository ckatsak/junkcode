# -*- coding: utf-8 -*-
"""

"""

import logging
import errno
import os
import time


FIFO1 = './one.fifo'

UNBUFFERED = 0
LINE_BUFFERED = 1

SLEEP_SECS = 0.8
JOB_SECS = 5
COMPUTE_SECS = 4
COMPUTE_TIMES = 3


def mkfifoz():
    try:
        os.mkfifo(FIFO1)
    except OSError as err:
        if err.errno != errno.EEXIST:
            logging.critical(err)
            raise


def sleep():
    time.sleep(SLEEP_SECS)
