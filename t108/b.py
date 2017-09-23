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

import comm


def job(num):
    for i in xrange(comm.JOB_SECS, 0, -1):
        logging.info("Job '%s': %d...", num, i)
        comm.sleep()


def main():
    comm.mkfifoz()

    logging.info("Opening '%s' for reading...", comm.FIFO1)
    fifo = open(comm.FIFO1, 'r')
    # fifo = open(comm.FIFO1, 'r', comm.UNBUFFERED)
    # fifo = open(comm.FIFO1, 'r', comm.LINE_BUFFERED)
    logging.info("Just opened '%s' for reading!", comm.FIFO1)

    while True:
        logging.info("Reading from '%s'...", comm.FIFO1)
        x = fifo.readline().strip()
        logging.info("Just read '%s' from '%s'...", x, comm.FIFO1)
        if x is None or x == '':  # EOF
            logging.info("Received EOF and reopen('r')ing '%s'...", comm.FIFO1)
            fifo = open(comm.FIFO1, 'r')
            logging.info("Just reopened '%s' for reading!", comm.FIFO1)
        else:
            job(x)


if __name__ == '__main__':
    logging.getLogger().setLevel('DEBUG')
    main()
