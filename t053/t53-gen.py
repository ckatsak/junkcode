#!/usr/bin/env python3
#
# t53-gen.py
#
# NOTE: Each time this script is run, ./t53-timed.h gets overwritten!
#
# Create a C header file for t53-timed.c, containing an array of unsigned long
# long int, named rands, to be used for benchmarking.
#
# Author: ckatsak
#
import random

header_file = "t53-timed.h"
array_size  = 10000
range64_min = 1 << 63
range64_max = (1 << 64) - 1

def main():
    with open(header_file, 'w') as fout:
        print("#ifndef __T53_H__\n#define __T53_H__", file=fout)
        print("\n#define ARRAY_SIZE\t{}".format(array_size), file=fout)
        print("\n#define RANGE64_MIN\t{}ULL\t/* 1<<63 */".format(range64_min), file=fout)
        print("#define RANGE64_MAX\t{}ULL\t/* 1<<64 - 1*/".format(range64_max), file=fout)
        print("\n// Ignore the useless warning, more info at " +
			"https://gcc.gnu.org/bugzilla/show_bug.cgi?id=45977".format(), file=fout)
        print("extern unsigned long long rands[] = {", file=fout)
        for i in range(array_size // 4):
            print("\t" + " {}ULL,".format(random.randint(range64_min, range64_max)) * 4, file=fout)
#            print("\t" + " {}ULL,".format(random.getrandbits(64)) * 4, file=fout)
        print("};\n\n#endif /* __T53_H__ */", file=fout)

if __name__ == "__main__":
    main()
