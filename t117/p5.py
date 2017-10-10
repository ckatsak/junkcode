#!/usr/bin/env python

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


# MODIFY DEPLOYMENT

try:
    del depl.metadata.annotations['CHRISTOULAS']
    print '\'CHRISTOULAS\' annotation removed locally\n\n\n'
except KeyError:
    print '\'CHRISTOULAS\' annotation not found locally\n\n\n'


# UPDATE DEPLOYMENT:

new_depl = cl.replace_namespaced_deployment(
        name='pynsider',
        namespace='default',
        body=depl,
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
