# -*- coding: utf-8 -*-
"""
    Overriding __deepcopy__ method to be used by copy.deepcopy for user-defined
    class with unpicklable objects, using a parent class without these objects.
"""

import copy
from threading import Lock


class A(object):
    def __init__(self, num):
        self.a = num
        self.b = 'lala'
        self.c = [self.a] * 5

    def __repr__(self):
        return '<%s: (%d, %s) --> %s>' % (
                self.__class__, self.a, self.b, self.c)


class B(A):
    def __init__(self, num):
        self.lock = Lock()
        super(B, self).__init__(num)

    def __deepcopy__(self, memo):
        return type(self)(copy.deepcopy(self.a, memo))
