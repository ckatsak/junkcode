#!/usr/bin/env python

import logging

from pkg.a import A
from pkg.b import B
from pkg.c import C


def main():
    logging.debug("Initializing objects A, B and C...")
    a = A(42)
    b = B('yo')
    c = C('master')
    logging.debug("Objects A, B and C are ready!")

    for i in range(3):
        logging.debug("round %d:", i)
        a.cry()
        b.cry()
        c.cry()


if __name__ == '__main__':
    logging.basicConfig(
        format='%(asctime)s %(name)-16s %(module)-5s %(levelname)-8s   %(message)s',
        level=logging.DEBUG)
    main()
