#!/usr/bin/env python

"""
    Client side for testing dynamic URL API endpoints, using python requests
    module.

    To run:

        $ pip install requests
        $ ./client.py dynamic_endpoint_FTW
"""

import logging
import sys

import requests

post_URL = 'http://127.0.0.1:55555/objects/'


def main():
    user_input = sys.argv[1] if len(sys.argv) > 1 else 'skata'
    logging.debug("user_input: %s" % user_input)

    try:
        # resp = requests.post(post_URL, json={'yoleles': user_input})
        # resp = requests.post(post_URL, data=user_input)
        resp = requests.post(post_URL, data=user_input, allow_redirects=False)
        logging.debug("HTTP POST Response: %r", resp)
    except Exception as e:
        logging.exception(e)
        sys.exit(42)

    # get_URL = resp.url
    get_URL = resp.headers['location']
    logging.info("New HTTP Redirect URL: %s", get_URL)

    try:
        resp = requests.get(get_URL)
        logging.debug("HTTP GET Response: %r", resp)
    except Exception as e:
        logging.exception(e)
        sys.exit(66)

    logging.info(resp.content)


if __name__ == '__main__':
    logging.basicConfig(level=logging.DEBUG)
    main()
