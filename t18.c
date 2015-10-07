#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <fcntl.h>
#include <sys/types.h>
#include <sys/wait.h>

int main(void)
{
//	fprintf(stderr, "[%d] My PID is %d\n\n", (int)getpid(), (int)getpid());
	fprintf(stderr, "[%s] My PID is %d\n\n", __FILE__, (int)getpid());

	return 0;
}
