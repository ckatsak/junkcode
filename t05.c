#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define SZ 20

int main(void)
{
	long int_p = 27, fract_p = 506;
	char buf_data[SZ];

	fprintf(stderr, "\n\tip = %ld\n\tfp = %ld\n\tbuf = %s\n\tstrlen(buf) ="
		" %zd\n\n", int_p, fract_p, buf_data, strlen(buf_data));

	sprintf(buf_data, "%ld.%ld", int_p, fract_p);

	fprintf(stderr, "\n\tip = %ld\n\tfp = %ld\n\tbuf = %s\n\tstrlen(buf) ="
		" %zd\n\n", int_p, fract_p, buf_data, strlen(buf_data));

	return 0;
}
