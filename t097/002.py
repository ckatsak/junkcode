#!/usr/bin/env python

import os
from pprint import pprint

from kubernetes import (
    client as klient,
    config as konfig,
    watch as kwatch,
)


def main():
    konfig.load_incluster_config()
    v1b1 = klient.AppsV1beta1Api()

    w = kwatch.Watch()
    try:
        for e in w.stream(v1b1.list_namespaced_deployment,
                          os.getenv('POD_NAMESPACE'),
                          label_selector='app=pyntainerdb',
                          timeout_seconds=180,
                          watch=True,
                          ):
            pprint(e)
    except klient.rest.ApiException as err:
        print err


if __name__ == '__main__':
    main()
