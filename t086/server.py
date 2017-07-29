#!/usr/bin/env python

"""
    Test dynamic URL API endpoint creation by means of HTTP redirect 303 SEE
    OTHER.

    - Client sends a POST request to /objects/. Request body contains a string
    with the name of the new URL endpoint (to be "created").
    - Server responds 303 SEE OTHER with /objects/list/<that_string_above> on
    its Location header.
    - Client can now parse the new URL from response's Location header, and
    issue a new GET request there.

    To run it:

        $ pip install Flask
        $ ./server.py

    Then try it one of these:

        $ curl -v -L -d 'dynamic_endpoints_FTW' http://localhost:55555/objects
        $ ./client.py dynamic_endpoints_FTW

    or using any other HTTP client of your choice.

"""

import json

from flask import Flask, redirect, request, url_for


app = Flask(__name__)


# # #  NO JSON CASE  # # #

@app.route('/objects/', methods=['POST'])
def list_objects():
    request.get_data()
    myvar = request.data
    return redirect(url_for('chunked_list', yoleles=myvar), code=303)


@app.route('/objects/list/<yoleles>', methods=['GET'])
def chunked_list(yoleles):
    return 'SERVER RESPONSE: %s' % yoleles


# # #  JSON CASE  # # #

@app.route('/objects/json/', methods=['POST'])
def list_objects_json():
    request.get_data()
    global myvar
    myvar = json.loads(request.data)['chunkSize']
    return redirect(url_for('chunked_list_json', yoleles=hash(myvar)),
                    code=303)


@app.route('/objects/list/json/<yoleles>', methods=['GET'])
def chunked_list_json(yoleles):
    global myvar
    return 'SERVER RESPONSE: %s' % myvar


if __name__ == '__main__':
    app.run(host='127.0.0.1',
            port=55555,
            threaded=True,
            debug=True)
