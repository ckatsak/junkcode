#!/usr/bin/env python

"""
    Test query parameters on Flask.
"""

from flask import Flask, request

app = Flask(__name__)


@app.route('/yoyoyo/', methods=['GET'])
def yoyoyo():
    app.logger.debug(request.args.get('skata', ''))
    if request.args.get('skata', ''):
        return 'YAY'
    else:
        return 'NAY'


if __name__ == '__main__':
    app.run(host='127.0.0.1', port=55555, debug=True)
