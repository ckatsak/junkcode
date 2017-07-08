#include <sys/types.h>
#include <sys/stat.h>
#include <unistd.h>
#include <stdio.h>


int main(void) {
	struct stat st;
	char *whichtty;

	if ( !isatty(STDIN_FILENO) ) {
		(void) fprintf(stderr, "%s\n", "Stdin is not connected to a tty device.");
	} else {
		if ( (whichtty = ttyname(STDIN_FILENO)) == NULL ) {
			perror("ttyname");
			return 1;
		} else {
			(void) fprintf(stderr, "Stdin is connected to tty device %s.\n", whichtty);
		}
	}

	if ( fstat(STDIN_FILENO, &st) == -1 ) {
		perror("stat");
		return 1;
	} else {
		if ( st.st_mode & S_IWUSR ) {
			(void) fprintf(stderr, "%s\n", "Stdin appears to be writeable by the user.");
		} else {
			(void) fprintf(stderr, "%s\n", "Stdin does not appear to be writeable by the user.");
		}
	}

	return 0;
}

