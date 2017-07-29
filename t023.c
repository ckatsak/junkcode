#include <stdio.h>
#include <stdlib.h>
#include <sys/time.h>
#include <unistd.h>
#include <errno.h>

#define BLOCK_SIZE	16

static const long long SIZE = 16384;

inline int min(int x, int y)
{
	return x < y ? x : y;
}

int main(void)
{
//	int i, j, x, y, a[SIZE][SIZE], b[SIZE], c[SIZE];
	long long i, j, x, y, **a, *b, *c, *t;
	struct timeval ts, tf;

	a = malloc(SIZE * sizeof(*a));
	b = malloc(SIZE * sizeof(*b));
	c = malloc(SIZE * sizeof(*c));
	t = malloc(SIZE * SIZE * sizeof(*t));
	if (!t || !a || !b || !c) {
		perror("malloc");
		exit(EXIT_FAILURE);
	}
	else {
		for (i = 0; i < SIZE; i++)
			a[i] = t + i * SIZE;
	}

	for (i = 0; i < SIZE; i++) {
		b[i] = 1;
		c[i] = 0;
		for (j = 0; j < SIZE; j++)
			a[i][j] = i + j + 1;
	}

	gettimeofday(&ts, NULL);

	for (i = 0; i < SIZE; i++)
		for (j = 0; j < SIZE; j++)
			c[i] += a[i][j] * b[j];

	gettimeofday(&tf, NULL);

	fprintf(stderr, "\n\tNaive time: %0.8lf\n\n", (double)((tf.tv_sec - ts.tv_sec) + (tf.tv_usec - ts.tv_usec) * 0.000001));

	gettimeofday(&ts, NULL);

	for (i = 0; i < SIZE; i += BLOCK_SIZE) {
		for (j = 0; j < SIZE; j += BLOCK_SIZE) {
			for (x = i; x < i + BLOCK_SIZE; x++)
				for (y = j; y < j + BLOCK_SIZE; y++) {
					c[x] += a[x][y] * b[y];
			}
		}
	}

	gettimeofday(&tf, NULL);

	fprintf(stderr, "\n\tNice time: %0.8lf\n\n", (double)((tf.tv_sec - ts.tv_sec) + 0.000001 * (tf.tv_usec - ts.tv_usec)));

	return 0;
}
