#include <stdio.h>
#include <stdlib.h>

#define WELCOME()							\
	do {								\
		fprintf(stderr, "\nEntering function: %s\n", __func__);	\
	} while(0)

#define GOODBYE()							\
	do {								\
		fprintf(stderr, "\nLeaving function: %s\n", __func__);	\
	} while(0)

static const int SIZE = 10;

void print2d(int *arr, int dimX, int dimY)
{
	WELCOME();
	int i, j;

	for (i = 0; i < dimX; i++) {
		for (j = 0; j < dimY; j++)
			fprintf(stderr, "\t%d", (arr + i * dimY)[j]);
		fprintf(stderr, "\n");
	}

	GOODBYE();
}

void print1d(int *arr, int len)
{
	WELCOME();
	int i;

	for (i = 0; i < len; i++)
		fprintf(stderr, "\t%d", arr[i]);

	GOODBYE();
}

void init2d(int *arr, int dimX, int dimY)
{
	WELCOME();
	int i, j;

	for (i = 0; i < dimX; i++)
		for (j = 0; j < dimY; j++)
			(arr + i * dimY)[j] = i + j;

	GOODBYE();
}

int main(void)
{
	int arr[SIZE][SIZE];

	init2d(*arr, SIZE, SIZE);
	print2d(*arr, SIZE, SIZE);
	print1d(*arr, SIZE * SIZE);

	return 0;
}
