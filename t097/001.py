#!/usr/bin/env python

import os
from pprint import pprint

from kubernetes import (
    client as klient,
    config as konfig,
    watch as kwatch,
)


def main():
    '''WRONG because read_namespaced_deployment doesn't support 'watch'
    argument.'''
    konfig.load_incluster_config()
    v1b1 = klient.AppsV1beta1Api()

    w = kwatch.Watch()
    try:
        for e in w.stream(v1b1.read_namespaced_deployment,
                          'pynsider',
                          os.getenv('POD_NAMESPACE'),
                          ):
            pprint(e)
    except klient.rest.ApiException as err:
        print err


if __name__ == '__main__':
    main()
