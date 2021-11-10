/*
 * Playing with THP over PMEM-DAX with ext4-dax
 *
 * Spawn it, `pmap -X <PID>` it and check its `/proc/<PID>/smaps`, in different
 * phases of its functionality (defined by you, via `kill -CONT <PID>`)
 */

#include <stdlib.h>
#include <unistd.h>
#include <stdio.h>
#include <signal.h>
#include <sys/mman.h>
#include <string.h>
#include <fcntl.h>
#include <sys/types.h>
#include <sys/stat.h>

#define TWO_MiB	(1 << 21)

int main(void)
{
	fprintf(stderr, "My PID: %d\n", getpid());
	fprintf(stderr, "Waiting for a SIGCONT ...\n\n");
	raise(SIGTSTP);

	//char *buf = NULL;
	//fprintf(stderr, "Allocating %d bytes...\n", TWO_MiB);
	//buf = mmap(NULL, TWO_MiB, PROT_READ | PROT_WRITE | PROT_EXEC,
	//		MAP_PRIVATE | MAP_ANONYMOUS, -1, 0);
	//if (MAP_FAILED == buf) {
	//	perror("mmap");
	//	return 1;
	//}
	//fprintf(stderr, "Waiting for a SIGCONT ...\n\n");
	//raise(SIGTSTP);
	//
	//fprintf(stderr, "Setting them to '0xff'...\n");
	//memset(buf, 0xff, TWO_MiB);
	//fprintf(stderr, "Waiting for a SIGCONT ...\n\n");
	//raise(SIGTSTP);
	
	char *fbuf = NULL;
	int fd;
	fprintf(stderr, "mmap(2)ing PMEM file\n");
	if (-1 == (fd = open("/mnt/pmem0/" __FILE__, O_RDONLY))) {
		perror("open");
		return 1;
	}
	fbuf = mmap(NULL, TWO_MiB, PROT_READ, MAP_PRIVATE, fd, 0);
	if (MAP_FAILED == fbuf) {
		perror("mmap");
		return 1;
	}
	fprintf(stderr, "Waiting for a SIGCONT ...\n\n");
	raise(SIGTSTP);

	char *abuf;
	fprintf(stderr, "Copying mmap(2)'d file in DRAM buffer\n");
	abuf = mmap(NULL, 2 * TWO_MiB, PROT_READ | PROT_WRITE | PROT_EXEC,
			MAP_PRIVATE | MAP_ANONYMOUS, -1, 0);
	if (MAP_FAILED == abuf) {
		perror("mmap");
		return 1;
	}
	memcpy(abuf, fbuf, 256);
	memcpy(abuf + TWO_MiB, fbuf, 256);
	fprintf(stderr, "Waiting for a SIGCONT ...\n\n");
	raise(SIGTSTP);

	return 0;
}

