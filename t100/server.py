#!/usr/bin/env python
# -*- coding: utf-8 -*-
"""
    Testing flask + signal handling + threading + fifo
"""

import os
import signal
# import sys
import threading

from flask import Flask

FS_FIFO_NAME = './yoyo.fifo'

app = Flask(__name__)


@app.route('/ping', methods=['GET'])
def ping():
    return 'pong'


def term_handler(signal, frame):
    print '[%s]: A wild SIGTERM appears!' % (
            threading.current_thread().getName())
    with open(FS_FIFO_NAME, 'w') as fifo:
        print '[%s]: Writing 1 to FIFO' % threading.current_thread().getName()
        fifo.write('1')


def extra_thread_fn():
    with open(FS_FIFO_NAME, 'r') as fifo:
        print '[%s]: Reading from FIFO' % threading.current_thread().getName()
        fifo.read(1)
        print '[%s]: Exiting...' % threading.current_thread().getName()
        # sys.exit(42)  # Raises SystemExit, so it's the same as thread.exit()
        os._exit(42)


if __name__ == '__main__':
    threading.Thread(target=extra_thread_fn, name='extra_thread').start()

    print '[%s]: PID: %d' % (threading.current_thread().getName(), os.getpid())
    signal.signal(signal.SIGTERM, term_handler)

    app.run(host='127.0.0.1',
            port=50000,
            threaded=True,
            debug=False)  # debug=False to catch the signal
