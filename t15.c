/**
 * Testing:
 *  - divide-by-zero exceptions
 *  - compile-time macro definition
 *
 * Compile either as:
 *	gcc t15.c -o t15 -Wall -O3 -DCONST
 * or as:
 *	gcc t15.c -o t15 -Wall -O3 -DVAR
 */

#include <stdio.h>

int main(void)
{

#ifdef CONST
	fprintf(stderr, "\n\n\t%d\n\n\n", 5 / 0);
#endif

#ifdef VAR
	int x = 5, y = 1;
	fprintf(stderr, "\n\n\t%d\n\n\n", x / y);
#endif

	return 0;
}
