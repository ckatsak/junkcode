#include <stdio.h>
#include <stdlib.h>

int main(void)
{
	int i, j, mat[10][10], **arr;

	arr = malloc(10 * sizeof(*arr));
	for (i = 0; i < 10; i++)
		arr[i] = malloc(10 * sizeof(**arr));

	fprintf(stderr, "\n\t2d matrix:\n");
	for (i = 0; i < 10; i++) {
		for (j = 0; j < 10; j++)
			fprintf(stderr, "%p\t", &(mat[i][j]));
		fprintf(stderr, "\n");
	}

	fprintf(stderr, "\n\t**array\n");
	for (i = 0; i < 10; i++) {
		for (j = 0; j < 10; j++)
			fprintf(stderr, "%p\t", &(arr[i][j]));
		fprintf(stderr, "\n");
	}

	fprintf(stderr, "\n\n");
	return 0;
}
