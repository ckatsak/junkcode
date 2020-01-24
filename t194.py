#!/usr/bin/env python2
# -*- coding: utf-8 -*-
#
# Simple demo script to telnet into the read-only account of our PDU.

import getpass
import telnetlib


def main():
    host = raw_input("Telnet server: ") or '147.102.4.150'
    username = raw_input("Username: ")
    password = getpass.getpass()

    print "Attempting connection..."
    tn = telnetlib.Telnet(host)
    print "Connection successful."

    print 'Reading until "User Name : "...'
    tn.read_until('User Name : ')

    print 'Sending user name "%s"...' % username
    tn.write('%s\r\n' % username)
    print 'User name "%s" sent successfully.' % username

    print "Reading until 'Password  : '..."
    tn.read_until('Password  : ')
    print "Sending the password..."
    tn.write('%s\r\n' % password)
    print "Password sent successfully."

    print 'Reading untul "apc>"...'
    tn.read_until('apc>')
    print "Prompt read successfully."

    print 'Sending command "olReading 1 power"...'
    tn.write('olReading 1 power\r\n')
    print "Command sent successfully."

    print 'Sending command "bye"...'
    tn.write('bye\r\n')
    print "Command 'bye' sent successfully."

    print "Reading all..."
    print tn.read_all()


if __name__ == '__main__':
    main()
