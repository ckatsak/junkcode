#include <stdio.h>
#include <string.h>
#include <stdlib.h>

void func(char *string)
{
	printf("Entering func\n");

	int authenticated = 0;
	char buffer[4];

	strcpy(buffer, string);

	if (authenticated) {
		printf("AUTHENTICATION SUCCESSFUL\n");
	}

	printf("Leaving func\n");
}

int main(int argc, char *argv[])
{
	printf("Entering main\n");

	if ( argc != 2) {
		fprintf(stderr, "ERROR: Usage: %s <string>\n", argv[0]);
		exit(1);
	}

	char string[10];
	strcpy(string, argv[1]);

	func(string);

	printf("Leaving main\n");

	return 0;
}
