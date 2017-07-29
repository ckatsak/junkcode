/*
 * Testing likely(x) and unlikely(x), GCC extensions macros.
 *
 * Build with and without the macros
 * $ gcc t53.c -Wall -o t53-with -O0
 * $ gcc t53.c -Wall -o t53-without -O0 -DWITHOUT
 *
 * Now time both versions
 * $ time ./t53-with
 * $ time ./t53-without
 *
 * Build with and without the macros into assembly code, and inspect it, spot
 * the magic, if any
 * $ gcc t53.c -Wall -S -o t53-with.S -O0
 * $ gcc t53.c -Wall -S -o t53-without.S -O0 -DWITHOUT
 * $ diff -y t53-with{.S,out.S} | less
 */
#define _GNU_SOURCE

#include <stdio.h>
#include <stdlib.h>
#include <time.h>
#include <limits.h>

#ifndef WITHOUT
 #define likely(x)	__builtin_expect (!!(x), 1)
 #define unlikely(x)	__builtin_expect (!!(x), 0)
#else
 #define likely(x)	(x)
 #define unlikely(x)	(x)
#endif

int main(void)
{
	register long long i;
	srand(time(NULL));
	register long long x = rand();
	
	for (i = 0; i < RAND_MAX; i++)
		if (unlikely(i == x)) {
			fprintf(stderr, "i == x == %lld\n", i);
			break;
		}

	return EXIT_SUCCESS;
}
