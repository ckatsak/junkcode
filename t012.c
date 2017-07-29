#include <stdio.h>
#include <stdlib.h>

int main(void)
{
	int *x = malloc(sizeof(int));

	fprintf(stderr, "\n\tx = %d\n", *x);

	{
		*x = 128;
		fprintf(stderr, "\n\t*x = %d\n", *x);
	}

	fprintf(stderr, "\n\tx = %d\n", *x);

	return 0;
}
