#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int main(void)
{
	struct {
		int x;
		unsigned char string[];
	} skata;
	unsigned char string[] = "Hello, World!";

	fprintf(stderr, "\n\tstring[] = %s\n", string);

//	fprintf(stderr, "\n\tskata.string[] = %s\n", skata.string);
//	strcpy(skata.string, string);
//	fprintf(stderr, "\n\tskata.string[] = %s\n", skata.string);

	fprintf(stderr, "\n\tskata.string[] = %s\n", skata.string);
	memcpy(skata.string, string, sizeof(string));
	fprintf(stderr, "\n\tskata.string[] = %s\n", skata.string);

	fprintf(stderr, "\n\n");
	return 0;
}
