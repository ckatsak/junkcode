#!/usr/bin/env python
# -*- coding: utf-8 -*-

import logging

from flask import Flask


app = Flask(__name__)


@app.route('/ping', methods=['GET'])
def ping():
    logging.debug('Ping! (debug)')
    logging.info('Ping! (info)')
    logging.warning('Ping! (warning)')
    logging.error('Ping! (error)')
    logging.critical('Ping! (critical)')
    return 'pong'


@app.route('/hello', methods=['GET'])
def hello():
    logging.debug('Hello! (debug)')
    logging.info('Hello! (info)')
    logging.warning('Hello! (warning)')
    logging.error('Hello! (error)')
    logging.critical('Hello! (critical)')
    return 'yo yo'


if __name__ == '__main__':
    logging.basicConfig(
            format='%(asctime)s:%(levelname)s:%(threadName)s:%(message)s')
    app.run(host='127.0.0.1', port=55555, debug=False)
