#include <stdio.h>
#include <stdint.h>

int detect_endianness(void)
{
	union {
		char arr[4];
		int32_t num;
	} test;

	for (int i = 0; i < 4; i++)
		test.arr[i] = i + 1;

	//fprintf(stderr, "test.num == 0x%x\n", test.num);
	return test.num == 0x01020304 ? 1 : 0;
}

int main(void)
{
	fprintf(stderr, "\nEndianness detected: %s\n\n",
			detect_endianness() ? "BIG" : "LITTLE");

	return 0;
}
