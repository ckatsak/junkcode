#!/usr/bin/env python
# -*- coding: utf-8 -*-
"""
    Run:
        $ gunicorn --bind 0.0.0.0:6000 --workers 4 --worker-class gevent \
                --threads 4 wsgi:app
    Test:
        $ for i in $(seq 1 10000); do curl http://localhost:6000/echo/$i & \
                ; done
"""

import json
import logging

from flask import make_response, request


class Server(object):
    def __init__(self):
        self._content = 42
        self._error = 666

    def echo_post(self):
        resp = make_response()
        resp.content_type = 'application/json'
        resp.status_code = 202
        cont = request.get_data()
        resp.set_data(cont)
        logging.debug('Post echo: %s (debug)', cont)
        logging.info('Post echo: %s (info)', cont)
        logging.warning('Post echo: %s  (warning)', cont)
        logging.error('Post echo: %s (error)', cont)
        logging.critical('Post echo: %s (critical)', cont)
        return resp

    def echo(self, string):
        resp = make_response()
        resp.content_type = 'application/json'
        resp.status_code = 201
        resp.set_data(json.dumps(string))
        logging.debug('Echo: %s (debug)', string)
        logging.info('Echo: %s (info)', string)
        logging.warning('Echo: %s  (warning)', string)
        logging.error('Echo: %s (error)', string)
        logging.critical('Echo: %s (critical)', string)
        return resp

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
