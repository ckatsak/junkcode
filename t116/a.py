#!/usr/bin/env python

import logging
import time


LOGFILE = './file.log'


def main():
    logging.info('YO YO')
    logging.debug(time.asctime())


if __name__ == '__main__':
    logging.basicConfig(
            filename=LOGFILE,
            format='%(asctime)s:%(levelname)-8s:%(threadName)-12s:%(message)s',
            level=logging.DEBUG)
    main()
