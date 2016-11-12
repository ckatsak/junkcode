#
# FIXME: this Makefile gets the job done, but SUCKS
#
CC=gcc
CFLAGS=-Wall -O0

all: plain timed

plain: t53-plain.c
	$(CC) $^ $(CFLAGS) -DWITHOUT -o $@-without
	$(CC) $^ $(CFLAGS) -o $@-with
	$(CC) $^ $(CFLAGS) -S -g -DWITHOUT -o $@-without.S
	$(CC) $^ $(CFLAGS) -S -g -o $@-with.S

timed: t53-timed.c
	/usr/bin/python3 t53-gen.py
	$(CC) $^ $(CFLAGS) -DWITHOUT -o $@-without
	$(CC) $^ $(CFLAGS) -o $@-with

clean:
	rm -vf t53-timed.h plain-with{,out} timed-with{,out}

.PHONY: all clean
