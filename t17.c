#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <fcntl.h>
#include <sys/types.h>
#include <sys/wait.h>

int main(void)
{
	pid_t pid;

	fprintf(stderr, "[%s] My PID is %d\n\n", __FILE__, (int)getpid());

	pid = fork();

	if (pid == 0) {
		char *argv[] = { "./t18", NULL };
		char *envp[] = { NULL };

		fprintf(stderr, "[%s] My PID is %d\n\n", __FILE__, (int)getpid());
		execve(argv[0], argv, envp);
		perror("execve");
	}

	wait(NULL);

	return 0;
}
