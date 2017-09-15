#!/usr/bin/env python
# -*- coding: utf-8 -*-
"""

"""

import logging

from flask import Flask

from server import Server

app = Flask(__name__)


if __name__ == '__main__':
    s = Server()
    logging.basicConfig(
            format='%(asctime)s:%(levelname)-8s:%(threadName)-12s:%(message)s')
    logging.getLogger().setLevel('DEBUG')

    app.add_url_rule('/ping', 'ping', s.ping)
    app.add_url_rule('/hello', 'hello', s.hello)
    app.add_url_rule('/content', 'content', s.content)
    app.add_url_rule('/error', 'error', s.error)
    app.run(host='127.0.0.1', port=45454, threaded=True, debug=False)
