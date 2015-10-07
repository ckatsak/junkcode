#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <sys/types.h>
#include <sys/wait.h>

int main(void)
{
	pid_t pid;
	int pfd[2], status;

	pipe(pfd);
	pid = fork();
	if (pid == 0) {
		char str1[100] = "skata";

		close(pfd[1]);
//		read(pfd[0], str1, 6);
		write(pfd[0], str1, 6);
		close(pfd[0]);
		printf("[%d]:\t%s\n", (int) getpid(), str1);
		exit(0);
	}
	else {
		while (wait(&status) != pid);
		char str[] = "hello";
		close(pfd[0]);
		write(pfd[1], str, strlen(str) + 1);
		close(pfd[1]);
		printf("[%d]:\t%s\n", (int) getpid(), str);
	}

	return 0;
}
