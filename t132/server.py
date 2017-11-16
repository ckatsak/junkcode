#!/usr/bin/env python
"""
    A simple (Flask) HTTP (forever-)streaming server of JSON data for tests.

    Quick client:
        $ curl -vN http://localhost:58080/stream
        $ curl -vN http://localhost:58080/range/9/20/2
        $ curl -vN http://localhost:58080/rand
    or:
        $ curl -vN http://localhost:58080/stream | jq '.' --unbuffered
        $ curl -vN http://localhost:58080/range/9/20/2 | jq '.' --unbuffered
        $ curl -vN http://localhost:58080/rand | jq '.' --unbuffered

"""

import json
import logging
import time

from flask import Flask, stream_with_context, Response


app = Flask('server')


@app.route('/stream', methods=['GET'])
def stream():
    @stream_with_context
    def _stream():
        i = 0
        l = []
        while True:
            i += 1
            l.append(hex(i)[2:])
            data = json.dumps(l)
            logging.info("Yielding %s", data)
            time.sleep(2)
            yield data
    return Response(_stream(), mimetype='application/json')


@app.route('/range/<int:start>/<int:end>/<int:step>', methods=['GET'])
def ranges(start, end, step):
    @stream_with_context
    def _ranges():
        prev, curr = start, start + step
        while curr < end:
            data = json.dumps(map(lambda i: hex(i)[2:], range(prev, curr)))
            logging.info("Yielding %s", data)
            yield data
            prev, curr = curr, curr + step
            time.sleep(2)
        if curr >= end:
            data = json.dumps(map(lambda i: hex(i)[2:], range(prev, end)))
            logging.info("Yielding %s", data)
            yield data
    return Response(_ranges(), mimetype='application/json')


@app.route('/rand', methods=['GET'])
def rand():
    t = int(time.time())
    if t % 2 == 0:
        logging.info("Allowed")
        return stream()
    else:
        logging.info("Disallowed")
        return Response("YOU SHALL NOT PASS",
                        status=418,
                        mimetype='application/json')


def main():
    logging.basicConfig(
            format='%(asctime)s %(threadName)16s %(levelname)8s %(message)s',
            level=logging.DEBUG)
    app.run(host='localhost', port=58080, debug=True, threaded=True)


if __name__ == '__main__':
    main()
