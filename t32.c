/*
 * Weird usage of #include - example.
 *
 * Doesn't need to be linked with "t32.h". Compiled with:
 *	$ gcc t32.c -o t32 -Wall -O3 -std=c99
 */
#include <stdio.h>
#include <stdlib.h>

#define N 10

int main(void)
{
	int x[N] = {
		#include "t32.h"
	};

	for (int i = 0; i < N; i++)
		fprintf(stderr, "\n\tx[%d] = %d\n", i, x[i]);

	return 0;
}
