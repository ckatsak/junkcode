#!/usr/bin/env python
"""
    It didn't work as expected.
"""

import logging
import sys

from kubernetes import client as klient, config as konfig
from kubernetes.client.rest import ApiException as kexc


logging.getLogger().setLevel('DEBUG')

konfig.load_incluster_config()
cl = klient.AppsV1beta1Api()


# READ DEPLOYMENT:

depl = cl.read_namespaced_deployment(
  name='pynsider',
  namespace='default',
  pretty='yeshh',
)
logging.info(depl.metadata)


# PATCH DEPLOYMENT:

new_md = depl.metadata
new_md.annotations['CHRISTOULAS'] = '42 re malakes'

new_depl = cl.patch_namespaced_deployment(
  name='pynsider',
  namespace='default',
  body=new_md,
  pretty='yeahhh',
)
logging.info(depl.metadata)
logging.info(new_depl.metadata)
