#include <stdio.h>
#include <stdlib.h>
#include <math.h>

int main(void)
{
	float x = 4.02f;
	x *= 100;
	int y = (int) x;
	fprintf(stderr, "x = %f\ty = %d\n\n", x, y);
}
