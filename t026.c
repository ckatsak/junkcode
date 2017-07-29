#include <stdio.h>
#include <limits.h>
#include <stdlib.h>
#include <omp.h>

int main(void)
{
	omp_set_num_threads(2);

	volatile int i;

	#pragma omp parallel
	{
		if (omp_get_thread_num() == 1) {
			fprintf(stderr, "ime h 1, tora arxizo\n");
			for (i = 0; i < INT_MAX; i++);
			fprintf(stderr, "ime h 1 k molis teliosa\n");
		}
		else
			fprintf(stderr, "ime h 0, tin kano apo edw\n");
	}

	fprintf(stderr, "ime h 0, tin ekana, h 1 akoma trexi\n");

	return 0;
}
