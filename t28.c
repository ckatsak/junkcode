#include <stdio.h>
#include "omp.h"

int main() {

  // omp_set_num_threads(NUM_THREADS);
  int i, k, p;
  int N=3;

  int A[3][3] = { {1, 2, 3},{ 5, 6, 7}, {8,9,10} };
  int B[3][3] =  { {1, 2, 3},{ 5, 6, 7}, {8,9,10} };
  int C[3][3] ;

  omp_set_dynamic(0);
   omp_set_num_threads(9);
   // printf("Num of threads %i \n", omp_get_max_threads());

	#pragma omp parallel for private(i,k,p) shared(A, B, C, N)
	for (p = 0; p < N * N; p++) {
		i = p / N;
		k = p % N;
           	int j = omp_get_thread_num();
       		C[i][k] = A[i][k] +  B[i][k] ;
              	printf("I m thread %d computing A[%d][%d] and B[%d][%d] = %d \n ", j, i,k, i,k, C[i][k]);
  	}

  int n, m;
  for (n=0; n<3; n++) {
    for ( m=0;m<3;m++){
      printf("C[%d][%d] = %d \n",n,m, C[n][m]);   

 }
}

return 0;

}
