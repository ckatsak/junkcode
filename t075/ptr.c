/*
 * Pointer assignment demo 
 */
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

struct str_t {
	int x, y;
};

typedef struct {
	int id;
	int *init;

	char buf_arr[32];
	char *buf_ptr;

	struct str_t s1;
	struct str_t *s2;
} request_t;

void pprint_req(char *name, request_t *req) {
	printf("\nRequest %s:\n", name);
	printf("\taddress: %p\n", (void *)req);

	printf("\tid: %d\n", req->id);
	printf("\tinit: %p\n", (void *)req->init);
	printf("\t  *init: %d\n", *req->init);

	printf("\tbuf_arr: %p\n", (void *)req->buf_arr);
	printf("\t  *buff_arr: \"%s\"\n", req->buf_arr);

	printf("\tbuf_ptr: %p\n", (void *)req->buf_ptr);
	printf("\t  *buf_ptr: \"%s\"\n", req->buf_ptr);

	printf("\ts1 (str_t):\n");
	printf("\t  req->s1.x: %d\n", req->s1.x);
	printf("\t  req->s1.y: %d\n", req->s1.y);

	printf("\ts2 (str_t *):\n");
	printf("\t  &s2: %p\n", (void *)req->s2);
	if (req->s2 != NULL) {
		printf("\t  req->s2->x: %d\n", req->s2->x);
		printf("\t  req->s2->y: %d\n", req->s2->y);
	}

	printf("\n");
	fflush(stdout);
}

int main(void) {
	request_t *a = malloc(sizeof(request_t));
	memset(a, 0, sizeof(request_t));
	strcpy(a->buf_arr, "Yo yo yo, world!");
	a->id = 42;
	a->init = malloc(sizeof(int));
	*a->init = 505;
	a->s1.x = a->s1.y = 17;
	pprint_req("a", a);

	request_t *b = a;
	pprint_req("b (b=a)", b);

	request_t *c = malloc(sizeof(request_t));
	memset(c, 0, sizeof(request_t));
	*c = *a;
	pprint_req("c (*c=*a)", c);

	printf("-------------------------------------------------------------\n");
	printf("Modifying a->buf_arr content...\n");
	strcpy(a->buf_arr, "Modified content!");
	pprint_req("a", a);
	pprint_req("b (b=a)", b);
	pprint_req("c (*c=*a)", c);

	printf("-------------------------------------------------------------\n");
	printf("Allocating a->buf_ptr content...\n");
	a->buf_ptr = malloc(32);
	strcpy(a->buf_ptr, "New a->buf_ptr content!");
	pprint_req("a", a);
	pprint_req("b (b=a)", b);
	pprint_req("c (*c=*a)", c);

	printf("-------------------------------------------------------------\n");
	printf("Allocating a->s2 content...\n");
	a->s2 = malloc(sizeof(struct str_t));
	a->s2->x = 1;
	a->s2->y = 2;
	pprint_req("a", a);
	pprint_req("b (b=a)", b);
	pprint_req("c (*c=*a)", c);

	printf("-------------------------------------------------------------\n");
	printf("Modifying *a->init content...\n");
	*a->init = 333;
	pprint_req("a", a);
	pprint_req("b (b=a)", b);
	pprint_req("c (*c=*a)", c);

	return 0;
}
