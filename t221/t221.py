#!/usr/bin/env python3

import json


PL = [
    123,
    34,
    110,
    114,
    111,
    119,
    34,
    58,
    32,
    55,
    54,
    44,
    32,
    34,
    110,
    99,
    111,
    108,
    34,
    58,
    32,
    55,
    54,
    125,
]

IN = {"nrow": 76, "ncol": 76}


def payloadize(d):
    return list(json.dumps(d).encode())


_ = json.dumps(IN)
print(f"json.dumps(IN) == {_}\nlen == {len(_)}\n")

_ = json.dumps(IN).encode()
print(f"json.dumps(IN).encode() == {_}\nlen == {len(_)}\n")

_ = list(json.dumps(IN).encode())
print(f"list(json.dumps(IN).encode()) == {_}\nlen == {len(_)}\n")

print(f"\nPL == {PL}\nlen == {len(PL)}\n")

_ = bytes(PL)
print(f"bytes(PL) == {_}\nlen == {len(_)}\n")

_ = json.loads(bytes(PL))
print(f"json.loads(bytes(PL)) == {_}\nlen == {len(_)}")
