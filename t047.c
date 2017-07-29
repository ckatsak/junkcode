/*
 * Testing ulimit(1)
 *
 * Open a shell and type in:
 *	$ ulimit -v 1000000
 *	$ ./this_exec
 * and watch it crash due to virtual memory limit.
 */
#include <stdio.h>
#include <stdlib.h>

#define SIZE 1000

int main(void)
{
	int i = 0;
	char *arr[SIZE];

	for (i = 0; i < SIZE; ++i) {
		fprintf(stderr, "%d. Bytes allocated so far: %d\n",
			i, i * 1 << 20);
		if ((arr[i] = malloc(1 << 20) /*1MB*/) == NULL) {
			fprintf(stderr, "\n%d. Couldn't allocate %d more bytes!\n",
				i, 1 << 20);
			exit(1);
		}
	}

	return 0;
}
