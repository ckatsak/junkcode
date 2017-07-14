#!/usr/bin/env python

from flask import Flask


app = Flask(__name__)


@app.route('/objects/<obj_hash>', methods=['GET'])
def read(obj_hash):
    return 'IN READ'


@app.route('/objects/list/stream', methods=['GET'])
def list():
    return 'IN LIST'


if __name__ == '__main__':
    app.run(host='127.0.0.1', port=55555, debug=True, threaded=True)
