#include <stdio.h>

int main(void)
{
	unsigned int minor = 50;
	unsigned int sid = minor >> 3;
	unsigned msr = minor - (sid << 3);

	printf("\n\tminor = %u\n\tsid = %u\n\tmsr = %u\n\n", minor, sid, msr);

	return 0;
}
