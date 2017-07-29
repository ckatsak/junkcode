#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <string.h>

int main(void)
{
	char buf[256];

	memset(buf, 0, sizeof(buf));

	read(0, buf, sizeof(buf));
	buf[sizeof(buf) - 2] = '!';
//	buf[sizeof(buf) - 1] = '\0';
	buf[sizeof(buf) - 1] = 'a';

	write(1, buf, sizeof(buf));

/*
	fprintf(stderr, "\n\tbuf = %s\n\n", buf);

	int i;
	for (i = 0; i < sizeof(buf); i++) {
		if (buf[i] == '\0')
			fprintf(stderr, "\n\t%d: '\\0'", i);
		if (buf[i] == '\n')
			fprintf(stderr, "\n\t%d: '\\n', %c", i, buf[i]);
		if (buf[i] != '\0' && buf[i] != '\n')
			fprintf(stderr, "\n\t%d: %c,", i, buf[i]);
	}
*/

	fprintf(stderr, "\n\n");
	return 0;
}
