#!/usr/bin/env python

import socket

#create socket
s = socket.socket()

#listen without bind
s.listen(10)

#print the random port number we got
print s.getsockname()

s.close()
