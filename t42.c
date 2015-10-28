/*
 * Simple fork/wait and pipe example.
 */
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <sys/wait.h>

int main(void)
{
	int pfd[2], st;
	char string[64];
	pid_t pid;

	fprintf(stdout, "[%d]\tSPAWNED\n", (int) getpid());

	fprintf(stdout, "[%d]\tCREATING PIPE\n", (int) getpid());
	sleep(2);
	if (pipe(pfd) < 0) {
		perror("pipe");
		exit(1);
	}

	fprintf(stdout, "[%d]\tFORKING\n", (int) getpid());
	sleep(2);
	if ((pid = fork()) < 0) {
		perror("fork");
		exit(1);
	}

	/* Child process */
	if (pid == 0) {
		fprintf(stdout, "[%d]\tSPAWNED\n", (int) getpid());
		sleep(2);

		fprintf(stdout, "[%d]\tSTRCPY\n", (int) getpid());
		sleep(2);
		memset(string, 0, sizeof(string));
		strncpy(string, "skata patera\n", strlen("skata patera\n"));


		fprintf(stdout, "[%d]\tCLOSING READING END\n", (int) getpid());
		sleep(2);
		if (close(pfd[0]) < 0) {
			perror("close");
			exit(1);
		}

		fprintf(stdout, "[%d]\tWRITING\n", (int) getpid());
		sleep(2);
		if (write(pfd[1], string, sizeof(string)) < 0) {
			perror("write");
			exit(1);
		}

		fprintf(stdout, "[%d]\tCLOSING WRITING END\n", (int) getpid());
		sleep(2);
		if (close(pfd[1]) < 0) {
			perror("close");
			exit(1);
		}

		fprintf(stdout, "[%d]\tEXITING . . .\n", (int) getpid());
		sleep(2);

		return 0;
	}

	/* Parent process */

	fprintf(stdout, "[%d]\tCLOSING WRITING END\n", (int) getpid());
	sleep(2);
	if (close(pfd[1]) < 0) {
		perror("close");
		exit(1);
	}

	fprintf(stdout, "[%d]\tREADING. . .\n", (int) getpid());
	sleep(2);
	if (read(pfd[0], string, sizeof(string)) < 0) {
		perror("read");
		exit(1);
	}

	fprintf(stdout, "[%d]\tWRITING. . .\n", (int) getpid());
	sleep(2);
	if (write(2, string, sizeof(string)) < 0) {
		perror("write");
		exit(1);
	}

	fprintf(stdout, "[%d]\tREAPING. . .\n", (int) getpid());
	sleep(2);
	if (waitpid(-1, &st, 0) < 0) {
		perror("waitpid");
		exit(1);
	}

	fprintf(stdout, "[%d]\tCLOSING READING END\n", (int) getpid());
	sleep(2);
	if (close(pfd[0]) < 0) {
		perror("close");
		exit(1);
	}

	fprintf(stdout, "[%d]\tEXITING . . .\n", (int) getpid());
	sleep(2);

	return 0;
}
