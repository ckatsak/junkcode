#!/usr/bin/env python

import time


class Main():
    def __init__(self):
        def nested():
            return time.time()
        self.f = nested

    def __call__(self):
        print time.time()
        time.sleep(1)
        print (self.f)()
        time.sleep(1)
        print time.time()


if __name__ == '__main__':
    Main()()
