/*
 * 2d array (non-contiguous) allocation 
 */
#include <stdio.h>
#include <stdlib.h>

void alloc2d(int ***);
int **alloc2d2(void);

int main(void)
{
	int i, j, **arr;

	alloc2d(&arr);
	//arr = alloc2d2();

	for (i = 0; i < 10; i++) {
		for (j = 0; j < 5; j++)
			fprintf(stderr, " %d", arr[i][j]);
		fprintf(stderr, "\n");
	}

	// Memory locations
	fprintf(stderr, "\n---------------------------------------------\n\n");
	for (i = 0; i < 10; i++) {
		for (j = 0; j < 5; j++)
			fprintf(stderr, " %p", (void *)(&arr[i][j]));
		fprintf(stderr, "\n");
	}

	return 0;
}

void alloc2d(int ***a)
{
	int i;

	*a = malloc(10 * sizeof(int*));
	for (i = 0; i < 10; i++)
		(*a)[i] = malloc(5 * sizeof(int));
}

int **alloc2d2(void)
{
	int i, **arr;

	arr = malloc(10 * sizeof(int *));
	for (i = 0; i < 10; i++)
		arr[i] = malloc(5 * sizeof(int));

	return arr;
}
