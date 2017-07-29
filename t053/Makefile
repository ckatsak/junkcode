#
# FIXME: this Makefile gets the job done, but SUCKS
#
CC=gcc
CFLAGS=-Wall -O0

all: plain timed

plain: t53-plain.c
	$(CC) $^ $(CFLAGS) -DWITHOUT -o $@-without
	$(CC) $^ $(CFLAGS) -o $@-with
	$(CC) $^ $(CFLAGS) -S -DWITHOUT -o $@-without.S
	$(CC) $^ $(CFLAGS) -S -o $@-with.S

timed: t53-timed.c
	/usr/bin/python3 t53-gen.py
	$(CC) $^ $(CFLAGS) -DWITHOUT -o $@-without
	$(CC) $^ $(CFLAGS) -o $@-with

clean:
	rm -vf t53-timed.h plain-with plain-without plain-with.S \
		plain-without.S timed-with timed-without

.PHONY: all clean
