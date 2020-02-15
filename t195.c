/*
 * Testing C's integral promotions and usual arithmetic conversions (ecp p29).
 */
#include <stdio.h>

int main(void)
{
	short s1, s2;
	fprintf(stderr, "sizeof(s1) == %zd\nsizeof(s2) == %zd\n"
			"sizeof((s1 + s2)) == %zd\n\n",
			sizeof(s1), sizeof(s2), sizeof((s1 + s2)));

	float f1, f2;
	fprintf(stderr,	"sizeof(f1) == %zd\nsizeof(f2) == %zd\n"
			"sizeof((f1 + f2)) == %zd\n\n",
			sizeof(f1), sizeof(f2), sizeof((f1 + f2)));

	fprintf(stderr, "sizeof(s1) == %zd\nsizeof(f2) == %zd\n"
			"sizeof((s1 + f2)) == %zd\n\n",
			sizeof(s1), sizeof(f2), sizeof((s1 + f2)));

	double d1;
	fprintf(stderr, "sizeof(d1) == %zd\nsizeof(f2) == %zd\n"
			"sizeof((d1 + f2)) == %zd\n\n",
			sizeof(d1), sizeof(f2), sizeof((d1 + f2)));

	char c1;
	fprintf(stderr, "sizeof(c1) == %zd\nsizeof(s2) == %zd\n"
			"sizeof((c1 + s2)) == %zd\n\n",
			sizeof(c1), sizeof(s2), sizeof((c1 + s2)));

	unsigned long ul1;
	long double ld1;
	fprintf(stderr, "sizeof(ul1) == %zd\nsizeof(ld1) == %zd\n"
			"sizeof((ul1 + ld1)) == %zd\n\n",
			sizeof(ul1), sizeof(ld1), sizeof((ul1 + ld1)));

	long l1;
	unsigned u1;
	fprintf(stderr, "sizeof(l1) == %zd\nsizeof(u1) == %zd\n"
			"sizeof((l1 + u1)) == %zd\n\n",
			sizeof(l1), sizeof(u1), sizeof((l1 + u1)));


	return 0;
}
