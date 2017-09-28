#!/usr/bin/env python

import signal

import gevent


def run_forever():
    gevent.sleep(10000)


if __name__ == '__main__':
    gevent.signal(signal.SIGQUIT, gevent.kill)
    gl = gevent.spawn(run_forever)
    gl.join()
