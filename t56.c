/*
 *  memory_test_2.c
 *
 *  Author: ckatsak <ckatsak@gmail.com>
 */

#include <stdio.h>
#include <stdlib.h>
#include <errno.h>
#include <string.h>
#include <unistd.h>
#include <sys/mman.h>

#define handle_error(msg)							\
	do {									\
		printf("msg %s:%d\t%s\n", __func__, __LINE__, strerror(errno));	\
		exit(EXIT_FAILURE);						\
	} while (0)

int main(void)
{
	/* Initializations */
	char cmd[32];
	sprintf(cmd, "cat /proc/%ld/maps", (long)getpid());

	/* Print current memory mappings */
	system(cmd);

	/* Create a mapping for ("allocate") the page which includes address 0xDEADBEEF */
	void *ret = mmap((void *)0xDEADBEEF, sizeof(long), PROT_READ | PROT_WRITE,
			MAP_PRIVATE | MAP_ANONYMOUS, -1, 0);
	/* Return value checking */
	if (ret == MAP_FAILED)
		handle_error("mmap");
	if ( (long)ret != 0xDEADBEEF) {
		printf("\n\tmmap's return value wasn't 0xDEADBEEF, but 0x%lX instead!\n",
				(unsigned long)ret);
		/* Don't give up! It's just that PAGE_SIZE
		 * isn't an even divisor of 0xDEADBEEF */
		long page_size = sysconf(_SC_PAGESIZE);
		if (page_size == -1)
			handle_error(sysconf);
		printf("\tPage size = %ld, 0xDEADBEEF %% %ld == %ld\n\n", page_size,
				page_size, 0xDEADBEEF % page_size);
	}

	/* Store number 0xDEADBEEF at address 0xDEADBEEF */
	*(long *)0xDEADBEEF = 0xDEADBEEF;
	/* Could this fail after a successful mmap call above?
	 * Hint: mmap only gives kernel a hint! See man 2 mmap */

	/* Print current memory mappings again to spot any difference */
	system(cmd);

	return 0;
}
