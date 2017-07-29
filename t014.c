/**
 *  prctl(2)
 *  sprintf(3)
 **/
/*
 * Run it on a terminal, and then hit:
 *
 * 	ps -ef | egrep -e t14
 *
 * on another one to see teh zombiez.
 */

#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <errno.h>

#include <sys/prctl.h>
#include <sys/types.h>
#include <sys/wait.h>

#define SLEEP_SEC	30
#define PROC_NUM 	32

int main(void)
{
	pid_t pid;
	int i, status;
	char name[16];

	for (i = 0; i < PROC_NUM; ++i) {
		// fork attempt
		if ((pid = fork()) < 0) {
			perror("fork");
			exit(1);
		}

		/* Child */
		if (pid == 0) {
			sprintf(name, "t14_%03d", i);
			name[7] = '\0';

			if (prctl(PR_SET_NAME, name, 0, 0, 0) < 0) {
				perror("prctl");
				exit(1);
			}

			// Child exits without entering the loop.
			exit(0);
		}
	}

	/* Parent */
	sleep(SLEEP_SEC);

	for (i = 0; i < PROC_NUM; ++i) {
		pid = wait(&status);

		if (WIFEXITED(status))
			fprintf(stderr, "\n\t[%d]\tProcess %d exited with"
				" status %d.\n", (int) getpid(), pid,
				WEXITSTATUS(status));
	}

	return 0;
}
