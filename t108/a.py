#!/usr/bin/env python
# -*- coding: utf-8 -*-
"""
    Run:
        $ python a.py
        $ python b.py
    At some point send a SIGINT at a.py (the writer process), and watch b.py
    blocking to open the FIFO anew.
    Also, at some point send a SIGINT to b.py (the reader process), and watch
    a.py retrying the write infinitely.
"""

import logging
import os

import comm


def compute():
    for i in range(comm.COMPUTE_SECS, 0, -1):
        logging.info("Computing: %d...", i)
        comm.sleep()


def main():
    comm.mkfifoz()

    logging.info("Opening '%s' for writing...", comm.FIFO1)
    fifo = open(comm.FIFO1, 'w')
    # fifo = open(comm.FIFO1, 'w', comm.UNBUFFERED)
    # fifo = open(comm.FIFO1, 'w', comm.LINE_BUFFERED)
    logging.info("Just opened '%s' for writing!", comm.FIFO1)

    for i in range(comm.COMPUTE_TIMES):
        compute()
        while True:
            logging.info("Writing '%s' to '%s'...", str(i), comm.FIFO1)
            fifo.write('%s\n' % i)
            try:
                fifo.flush()
                logging.debug("\"try\" after flush")
            except IOError as err:
                logging.exception(err)
                comm.sleep()
                logging.debug("Continuing...")
                continue
            else:
                logging.debug("Break!")
                break
        logging.info("Just wrote '%s' to '%s'!", str(i), comm.FIFO1)
    os._exit(1)


if __name__ == '__main__':
    logging.getLogger().setLevel('DEBUG')
    main()
