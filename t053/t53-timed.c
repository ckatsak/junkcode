/*
 * Testing likely(x) and unlikely(x), GCC extensions macros.
 *
 * Build with and without the macros
 * $ gcc t53-timed.c -Wall -o t53-timed-with -O0
 * $ gcc t53-timed.c -Wall -o t53-timed-without -O0 -DWITHOUT
 * (Maybe also use flag -fno-guess-branch-probability ?)
 *
 * Now time both versions
 * $ time ./t53-timed-with
 * $ time ./t53-timed-without
 *
 * Build with and without the macros into assembly code, and inspect it, spot
 * the magic, if any
 * $ gcc t53-timed.c -Wall -S -o t53-timed-with.S -O0
 * $ gcc t53-timed.c -Wall -S -o t53-timed-without.S -O0 -DWITHOUT
 * $ diff -y t53-timed-with{.S,out.S} | less
 *
 * Author: ckatsak
 *
 * TODO: Benchmark with a proper random number generator function.
 * TODO: Makefile
 */
#define _GNU_SOURCE

#include <stdio.h>
#include <stdlib.h>
#include <limits.h>
#include <sys/time.h>

#include "t53-timed.h"

#ifndef WITHOUT
 #define likely(x)	__builtin_expect (!!(x), 1)
 #define unlikely(x)	__builtin_expect (!!(x), 0)
#else
 #define likely(x)	(x)
 #define unlikely(x)	(x)
#endif

int main(void)
{
	register unsigned long long i, j;
	struct timeval start, end;

	if (gettimeofday(&start, NULL)) {
		perror("gettimeofday(start)");
		return EXIT_FAILURE;
	}

	for (j = 0; j < ARRAY_SIZE; j++) 
		for (i = RANGE64_MIN; i < RANGE64_MAX - 1; i++)
			if (unlikely(i == rands[j]))
				break;

	if (gettimeofday(&end, NULL)) {
		perror("gettimeofday(end)");
		return EXIT_FAILURE;
	}

	fprintf(stderr, "\nTime elapsed: %ldusec.\n\n",
			(end.tv_sec - start.tv_sec) * 1000000 +
			(end.tv_usec - start.tv_usec));

	return EXIT_SUCCESS;
}
