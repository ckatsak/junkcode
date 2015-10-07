/*
 * Pi calculator - Monte Carlo method
 *
 * Compile using:
 *	$ gcc t33.c -o t33 -Wall -O3 -lm
 * Optional -DDEBUG for when debugging
 *
 * Author: Christos Katsakioris
 *	   (christosk24 [at] hotmail [dot] com)
 */
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <errno.h>
#include <fcntl.h>
#include <math.h>
#include <limits.h>
#include <omp.h>

static const long double r = 1.0L;

static long double
rand_ld(int urfd, long long *buf)
{
	int ret;

	ret = read(urfd, buf, sizeof(*buf));
	if (ret < 0) {
		perror("read(urandfd)");
		exit(EXIT_FAILURE);
	}

	return ((double) llabs(*buf)) / LLONG_MAX;
}

int main(int argc, char *argv[])
{
	int i, urandfd;
	long long REPS, buffer, quadrant, square;
	long double x, y;

	/* Parse argv[] */
	if (argc == 1)
		REPS = 30000;
	else if (argc == 2)
		REPS = atol(argv[1]);
	else {
		fprintf(stderr, "Usage: %s <REPS>\n", argv[0]);
		exit(EXIT_FAILURE);
	}
#ifdef DEBUG
	fprintf(stderr, "\n\tDEBUG: REPS = %lld\n", REPS);
	REPS = 3;
	fprintf(stderr, "\n\tDEBUG: REPS = %lld\n", REPS);
	fprintf(stderr, "\n\tDEBUG: omp_get_max_threads() = %d\n", omp_get_max_threads());
#endif

	/* Initialize /dev/urandom */
	urandfd = open("/dev/urandom", O_RDONLY);
	if (urandfd < 0) {
		perror("open(/dev/urandom)");
		exit(EXIT_FAILURE);
	}

	quadrant = square = 0;

	/* Loop and calculate ratio */
	for (i = 0; i < REPS; i++) {
		/* Roll the dice */
		x = rand_ld(urandfd, &buffer);
		y = rand_ld(urandfd, &buffer);
#ifdef DEBUG
		fprintf(stderr, "\n\tDEBUG: x = %Lf\n", x);
		fprintf(stderr, "\n\tDEBUG: y = %Lf\n", y);
#endif
		/* Update frequency */
		(sqrtl(x * x + y * y) < r) ? quadrant++ : square++;
/*		if (sqrtl(x * x + y * y) < r)
			quadrant++;
		else
			square++;
*/
#ifdef DEBUG
		fprintf(stderr, "\n\tDEBUG: --------------------------------------------------------\n");
#endif
	}

#ifdef DEBUG
	fprintf(stderr, "\n\tDEBUG: quadrant = %lld\n", quadrant);
	fprintf(stderr, "\n\tDEBUG: square = %lld\n", square);
#endif
	fprintf(stdout, "\n\n\tpi = %Lf\n\n",
		((double) quadrant / square) * r * r);

	return EXIT_SUCCESS;
}
