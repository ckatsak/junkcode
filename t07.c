#include <stdio.h>

int main(void)
{
	int x = 5;
	int *_x = &x;

	fprintf(stderr, "The address of x is: %p \n", &x);
	fprintf(stderr, "The value of _x is: %p\n", _x);

	return 0;
}
