#include <string.h>
#include <stdio.h>

int main(void)
{
	char bar[128];
	char *baz = &bar[0];

	baz[127] = 0;

	fprintf(stderr, "\nstrlen(baz) = %zd\n\n", strlen(baz));

	return (int)strlen(baz);
}
