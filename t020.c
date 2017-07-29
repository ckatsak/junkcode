#include <stdio.h>

int main(void)
{
	fprintf(stderr, "\n\n\tsizeof(long double) = %zd\n\n",
		sizeof(long double));

	return 0;
}
