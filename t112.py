import gevent
from gevent import getcurrent
from gevent.pool import Group

group = Group()


# map

def hello_from(n):
    print('[%d]: Size of group %s' % (n, len(group)))
    print('[%d]: Hello from Greenlet %x' % (n, id(getcurrent())))


group.map(hello_from, xrange(3))


# imap

def intensive(n):
    gevent.sleep(3 - n)
    return 'task', n


print('\nOrdered')


ogroup = Group()
for i in ogroup.imap(intensive, xrange(3)):
    print(i)


# imap_unordered

print('\nUnordered')

igroup = Group()
for i in igroup.imap_unordered(intensive, xrange(3)):
    print(i)
