#!/usr/bin/env python
"""
    This one actually works :D

    Add an annotation to a deployment (tested from a pod of that deployment).
"""

import logging
from pprint import pprint
import time

from kubernetes import client as klient, config as konfig


logging.getLogger().setLevel('DEBUG')

konfig.load_incluster_config()
cl = klient.AppsV1beta1Api()


# READ DEPLOYMENT:

depl = cl.read_namespaced_deployment(
  name='pynsider',
  namespace='default',
  pretty='yeshh',
)
print 'READ 1:\n'
pprint(depl.metadata)
print '\n\n\n'


# UPDATE DEPLOYMENT:

nd = klient.AppsV1beta1Deployment()
nd.api_version = depl.api_version
nd.kind = depl.kind
nd.metadata = depl.metadata
nd.metadata.annotations['CHRISTOULAS'] = '42 re malakes'
nd.spec = depl.spec

new_depl = cl.replace_namespaced_deployment(
  name='pynsider',
  namespace='default',
  body=nd,
  pretty='yup',
)
print 'RETURNED:\n'
pprint(new_depl)
print '\n\n\n'


time.sleep(10)


# READ DEPLOYMENT:

depl = cl.read_namespaced_deployment(
  name='pynsider',
  namespace='default',
  pretty='yeshh',
)
print 'READ 2:\n'
pprint(depl.metadata)
print '\n\n\n'
