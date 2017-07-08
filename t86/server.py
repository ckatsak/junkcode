#!/usr/bin/env python

"""
    Test dynamic URL creation by means of HTTP redirect 303 SEE OTHER.

    - Client sends a POST request to /objects/. Request body contains a string
    with the name of the new URL endpoint (to be "created").
    - Server responds 303 SEE OTHER with /objects/list/<that_string_above> on
    its Location header.
    - Client can now parse the new URL from response's Location header, and
    issue a new GET request there.

    Run:

        $ pip install Flask
        $ ./t86.py
        $ curl -v -L -d 'dynamic_endpoints_FTW' http://localhost:55555/objects

"""

from flask import Flask, redirect, request, url_for


app = Flask(__name__)


@app.route('/objects/', methods=['POST'])
def list_objects():
    request.get_data()
    myvar = request.data
    return redirect(url_for('chunked_list', yoleles=myvar), code=303)


@app.route('/objects/list/<yoleles>', methods=['GET'])
def chunked_list(yoleles):
    return 'SERVER RESPONSE: %s' % yoleles


if __name__ == '__main__':
    app.run(host='127.0.0.1',
            port=55555,
            threaded=True,
            debug=True)
