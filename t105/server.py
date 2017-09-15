#!/usr/bin/env python
# -*- coding: utf-8 -*-
"""

"""

import logging


class Server(object):
    def __init__(self):
        self._content = 42
        self._error = 666

    def ping(self):
        logging.debug('Ping! (debug)')
        logging.info('Ping! (info)')
        logging.warning('Ping! (warning)')
        logging.error('Ping! (error)')
        logging.critical('Ping! (critical)')
        return 'pong'

    def hello(self):
        logging.debug('Hello! (debug)')
        logging.info('Hello! (info)')
        logging.warning('Hello! (warning)')
        logging.error('Hello! (error)')
        logging.critical('Hello! (critical)')
        return 'yo yo'

    def content(self):
        logging.debug('Content! (debug)')
        logging.info('Content! (info)')
        logging.warning('Content! (warning)')
        logging.error('Content! (error)')
        logging.critical('Content! (critical)')
        return repr(self._content)

    def error(self):
        logging.debug('Error! (debug)')
        logging.info('Error! (info)')
        logging.warning('Error! (warning)')
        logging.error('Error! (error)')
        logging.critical('Error! (critical)')
        return repr(self._error)
