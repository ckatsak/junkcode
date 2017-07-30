#!/usr/bin/env python


class Foo(object):
    def __init__(self):
        self._greeting = 'Hello, world!'

    def __call__(self):
        print(self._greeting)


Foo()()
