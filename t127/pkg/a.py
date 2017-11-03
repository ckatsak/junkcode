import logging


class A(object):
    def __init__(self, x):
        self.x = x
        self._logger = logging.getLogger(__name__)
        self._logger.debug("ALL SET!")

    def cry(self):
        self._logger.debug("%s: %s", type(self).__name__, self.x)
