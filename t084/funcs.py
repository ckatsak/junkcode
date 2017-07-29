#!/usr/bin/env python

from flask import Flask

app = Flask(__name__)


@app.route('/show/<int:num>', methods=['GET'])
def show(num):
    return str(num)


@ app.route('/inc1/<int:num>', methods=['GET'])
def inc1(num):
    return str(num+1)


if __name__ == '__main__':
    app.run(host='0.0.0.0', debug=True)
