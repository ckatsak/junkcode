#!/usr/bin/env python

'''
    Test flask's `path` URL variable type.
'''

from flask import Flask

app = Flask(__name__)


@app.route('/<path:path>', methods=['GET'])
def index(path):
    return path


if __name__ == '__main__':
    app.run(host='127.0.0.1', port=55555, debug=True)
