#!/usr/bin/env python

from flask import Flask
from flask.views import View

app = Flask(__name__)


class Index(View):
    methods = ['GET']

    def dispatch_request(self):
        return 'Yo, yo, yo!'


class ShowNum(View):
    methods = ['GET']

    def __init__(self, num):
        self.num = str(num)

    def dispatch_request(self):
        return self.num


# @app.route('/inc1/<int:num>', methods=['GET'])
def inc1(num):
    return str(num+1)


if __name__ == '__main__':
    app.add_url_rule('/', view_func=Index.as_view('index'))
    app.add_url_rule('/show', view_func=ShowNum.as_view(
        'show_num', num=666))
    app.run(host='0.0.0.0', debug=True)
