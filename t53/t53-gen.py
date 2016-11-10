#!/usr/bin/env python3

import random

SIZE = 1000000

def main():
    """
    NOTE: Each time this script is run, ./t53.h gets overwritten!

    Create a C header file for t53.c, containing an array of long long int,
    named rands, to perform an accurate benchmark.
    """

    with open("t53.h", 'w') as fout:
        print("#ifndef __T53_H__", file=fout)
        print("#define __T53_H__", file=fout)
        print("\n#define SIZE\t{}".format(SIZE), file=fout)
        print("\nextern long long rands[] = {", file=fout)
        for i in range(SIZE // 4):
            print("\t{},\t{},\t{},\t{},".format(random.getrandbits(64),
                random.getrandbits(64), random.getrandbits(64),
                random.getrandbits(64)), file=fout)
        print("}\n\n#endif /* __T53_H__ */", file=fout)

if __name__ == "__main__":
    main()
