/*
 * $ gcc t49.c -o t49 -Wall -O3
 * $ sudo ./t49
 *
 * Watch it change its root directory to /home/christos
 */
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>

int main(void)
{
	char cwd[1024];

	// print current working directory
	if (getcwd(cwd, sizeof(cwd)) == NULL) {
		perror("getcwd() error");
		return EXIT_FAILURE;
	}
	fprintf(stderr, "Current working dir: %s\n", cwd);

	// chdir(2) or chroot(2) can break later. TODO: Why?
	if (chdir("/home/christos") != 0) {
		perror("chdir");
		return EXIT_FAILURE;
	}

	// chroot(2)
	if (chroot("/home/christos") != 0) {
		perror("chroot /tmp");
		return EXIT_FAILURE;
	}

	// print current working directory
	if (getcwd(cwd, sizeof(cwd)) == NULL) {
		perror("getcwd() error");
		return EXIT_FAILURE;
	}
	fprintf(stderr, "Current working dir: %s\n", cwd);

	// go back to /home/christos/junkcode
	if (chdir("/junkcode") != 0) {
		perror("chdir");
		return EXIT_FAILURE;
	}

	// print current working directory
	if (getcwd(cwd, sizeof(cwd)) == NULL) {
		perror("getcwd() error");
		return EXIT_FAILURE;
	}
	fprintf(stderr, "Current working dir: %s\n", cwd);

	return EXIT_SUCCESS;
}
