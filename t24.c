#include <stdio.h>
#include <stdlib.h>
#include <time.h>

#define MAX	(32 * 1024 * 32)
#define REP	100
#define B	(16 * 1024)

int main(void)
{
	int i, j, r;
	char array[MAX];

	for (i = 0; i < MAX; i++)
		array[i] = 0;

	clock_t t1 = clock();

	// Tiled loop
	for (i = 0; i < MAX; i += B) {
		for (r = 0; r < REP; r++) {
			for (j = i; j < i + B; j += 64) {
				array[j] = r;
			}
		}
	}

	clock_t t2 = clock();

	// Un-tiled loop
	for (r = 0; r < REP; r++) {
		for (i = 0; i < MAX; i += 64) {
			array[i] = r;
		}
	}

	clock_t t3 = clock();

	fprintf(stderr, "\nTiled\t: %f sec\n", (double)(t2 - t1) / CLOCKS_PER_SEC);
	fprintf(stderr, "\nUntiled\t: %f sec\n", (double)(t3 - t2) / CLOCKS_PER_SEC);
	fprintf(stderr, "array[0] = %d\n", array[0]);

	return 0;
}
