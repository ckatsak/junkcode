#include <stdio.h>

int main(void)
{
	volatile int u = 42;

	fprintf(stderr, "(u > 0) - (u < 0) == %d - %d == %d\n",
		(u > 0), (u < 0), (u > 0) - (u < 0));

	fprintf(stderr, "\n!((u > 0) - (u < 0)) == !(%d - %d) == %d\n",
		(u > 0), (u < 0), !((u > 0) - (u < 0)));

	fprintf(stderr, "\n!0 == %d\n!1 == %d\n", !0, !1);

	return 0;
}
