import logging

from .d import D


class C(object):
    def __init__(self, x):
        self.x = x
        self._d = D('puppet')
        self._logger = logging.getLogger(__name__)
        self._logger.debug("ALL SET!")

    def cry(self):
        self._logger.debug("%s: %s", type(self).__name__, self.x)
        self._d.cry()
