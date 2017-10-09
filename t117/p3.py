#!/usr/bin/env python

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
