/*
 * t19.c
 *
 * strace ./t19
 *
 * ltrace -S ./t19
 *
 * ./t19 and strace -p {./t19's pid}
 */

#include <stdio.h>
#include <unistd.h>
#include <fcntl.h>
#include <sys/types.h>

int main(void)
{
	volatile int i = 0;

	int fd = open("/dev/null", O_WRONLY);

	printf("%d\n", (int) getpid());

	for (;; ++i)
		if ( i % 100000000 == 0)
			write(fd, &i, sizeof(i));

	return 0;
}
