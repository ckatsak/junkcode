/*
 * Familiarization with unions basics.
 */
#include <stdio.h>
#include <limits.h>
#include <string.h>

int main(void)
{
	union {
		int a;
		char b;
		long long c;
	} u;

	fprintf(stderr, "\nsizeof(u) == %zd bytes == %zd bits\n\n",
			sizeof(u), sizeof(u) * 8);

//	memset(&u, 0, sizeof(u));
	u.a = 42;
	fprintf(stderr, "u.a == %d\nu.b == %c\nu.c == %lld\n\n",
			u.a, u.b, u.c);

//	memset(&u, 0, sizeof(u));
	u.b = 'a';
	fprintf(stderr, "u.a == %d\nu.b == %c\nu.c == %lld\n\n",
			u.a, u.b, u.c);

//	memset(&u, 0, sizeof(u));
	u.c = LLONG_MAX;
	fprintf(stderr, "u.a == %d\nu.b == %c\nu.c == %lld\n\n",
			u.a, u.b, u.c);

	return 0;
}
