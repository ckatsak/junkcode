# -*- coding: utf-8 -*-
"""Testing Python's property objects."""


class Foo(object):
    def __init__(self, x=24):
        self._x = x

    @property
    def x(self):
        return self._x

    @x.setter
    def x(self, value):
        self._x = value

    @x.deleter
    def x(self):
        del self._x


class Bar(object):
    """COMPLETELY WRONG"""
    N = 5

    def __init__(self, x=24):
        self._list = [x] * Bar.N

    @property
    def lst(self, index):
        return self._list[index]

    @lst.setter
    def lst(self, index, value):
        self._list[index] = value


class Baz(object):
    """Example usage:

    >>> z = Baz()
    Get:
    >>> z.d
    >>> z.d['c']  # seen as 2 steps: a) property getter, b) index appl. on _d
    Set: Impossible (unless d.setter is implemented)
    >>> z.d = {}
    AttributeError: can't set attribute
    Del: Impossible (unless d.deleter is implemented)
    >>> del z.d
    AttributeError: can't delete attribute

    So:
     + we're allowed to get a reference to _d using d
     + we're allowed to modify _d's content in every normal way
     - we're NOT allowed to set _d to something else, such as {}
     - we're NOT allowed to del _d
    """
    def __init__(self):
        self._d = {'a': 0, 'b': 1, 'c': 2, 'd': 3, 'e': 4}

    @property
    def d(self):
        return self._d
