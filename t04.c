#include <stdio.h>
#include <stdlib.h>

enum en { first = 1, second, third };

int main(void)
{
	enum en myenum = first;

	printf("\n%d\n\n", myenum << 2);

	return 0;
}
