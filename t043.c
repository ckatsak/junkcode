/*
 * Testing const char* type between different functions.
 */
#include <stdio.h>
#include <unistd.h>

const char * ignition_status(int st)
{
    const char *ign_st;

    if(st==1) {
       ign_st="ON";
    } else {
       ign_st="OFF";
    }
	printf("%p: %s\n", ign_st, ign_st);

    return ign_st;
}

int main()
{
	const char *a, *b, *c, *d;
	const char* skata = ignition_status(1);

	printf("%s\n", ignition_status(1));
	printf("%p: %s\n", skata, skata);

	const char *e;
	printf("%p: %s\n", a, a);
	printf("%p: %s\n", b, b);
	printf("%p: %s\n", c, c);
	printf("%p: %s\n", d, d);
	printf("%p: %s\n", e, e);

    printf("%s\t\t%p\n", skata, skata);

	int A[4096000], i;
//	printf("%p\n", A);
	for (i = 0; i < 4096; ++i)
		A[i] = 9999;

//	printf("%s\n", skata);

    return 0;
}
