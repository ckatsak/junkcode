#include <stdio.h>
#include <limits.h>

int main(void)
{
	volatile long long i = LLONG_MIN;

	fprintf(stderr, "[>");
	for (; i < LLONG_MAX; i++)
		if (!(i % 1000000000))
			fprintf(stderr, "\b=>");
	fprintf(stderr, "]");

	return 0;
}
