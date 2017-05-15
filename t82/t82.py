#!/usr/bin/env python

from pprint import pprint

import srvlookup


# recs = srvlookup.lookup('xmpp-server', 'TCP', 'gmail.com')
# pprint(recs)

def _pods_in_shard():
    raw_recs = srvlookup._query_srv_records('_%s._%s.%s' % ('xmpp-server',
            'TCP', 'gmail.com' or srvlookup._get_domain()))

    recs = sorted(srvlookup._build_result_set(raw_recs),
                  key=lambda r: (r.priority, -r.weight))

    for rec in recs:
        yield (rec.host, rec.port)


for pod in _pods_in_shard():
    pprint(pod)
