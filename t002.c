#include <stdio.h>

int main(void)
{
	int x = 5;
	fprintf(stderr, "\n\tx = %d\n", x);

	typeof(x) y = x;
	fprintf(stderr, "\n\ty = %d\n", y);

	fprintf(stderr, "\n");
	return 0;
}
