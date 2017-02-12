/*
 * Simple pointer arithmetic demonstration.
 */
#include <stdio.h>

int main(void)
{
	int *i = NULL;
	double *d = NULL;
	struct {
		double d;
		int i;
		//int i2;
		//int i3;
	} *s = NULL;

	fprintf(stderr, "sizeof(*i) == %zd, sizeof(i) == %zd\n",
			sizeof(*i), sizeof(i));
	fprintf(stderr, "sizeof(*d) == %zd, sizeof(d) == %zd\n",
			sizeof(*d), sizeof(d));
	fprintf(stderr, "sizeof(*s) == %zd, sizeof(s) == %zd\n",
			sizeof(*s), sizeof(s));

	fprintf(stderr, "\n--------------------------------------------\n\n");

	fprintf(stderr, "\ti == %p, i+1 == %p\n", (void *)i, (void *)(i + 1));
	fprintf(stderr, "\t(i+1) - i == %ld\n\n",
			((char *)(i + 1) - (char *)i));

	fprintf(stderr, "\td == %p, d+1 == %p\n", (void *)d, (void *)(d + 1));
	fprintf(stderr, "\t(d+1) - d == %ld\n\n",
			((char *)(d + 1) - (char *)d));
	
	fprintf(stderr, "\ts == %p, s+1 == %p\n", (void *)s, (void *)(s + 1));
	fprintf(stderr, "\t(s+1) - s == %ld\n\n",
			((char *)(s + 1) - (char *)s));

	return 0;
}
