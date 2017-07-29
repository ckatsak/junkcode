#include <stdio.h>
#ifdef _OPENMP
#include <omp.h>
#define ID omp_get_thread_num()
#else
#define ID 0
#endif

double a[100];

int main(int argc, char *argv[])
{
    int i;
    double rho;

/**/omp_set_num_threads(4);

    #pragma omp parallel
    {
        #pragma omp for
        for (i = 0; i < 100; i++) {
            a[i] = 2;
        }
        // MARKER
        rho = 0.0;
        #pragma omp for reduction(+: rho)
        for (i = 0; i < 100; i++) {
            rho += ((a[i])*(a[i]));
        }
        fprintf(stderr, "[%d] rho=%f\n", ID, rho);
    }
    fprintf(stderr, "[%d] rho=%f\n", ID, rho);
    return 0;
}
