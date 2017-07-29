#!/usr/bin/env python

import os.path

INDEX_HEIGHT = 3


hex_digits = map(lambda x: str(hex(x))[2], range(16))


def next_subdir():
    # level_state = [hex_digits[0]] * INDEX_HEIGHT
    level_state = [-1] * (INDEX_HEIGHT - 1)
    final_state = [15] * (INDEX_HEIGHT - 1)
    parent = ['0'] * (INDEX_HEIGHT - 1)
    level = 0
    level_state[0] = 0
    while True:
        # An eimai ston patera twn fyllwn, dwse ola ta fylla
        if level == INDEX_HEIGHT - 2:
            # Dwse ola ta fylla
            # parent = os.path.join(*map(lambda x: str(hex(x))[2],level_state))
            for i in range(16):
                yield os.path.join(parent[level], hex_digits[i])
            if level_state == final_state:
                break
            # Proxwra ston patera mou (gia na se steilei ston dipla mou h panw)
            parent[level] = ''
            level -= 1
        else:
            # An teleiwsa me ta paidia mou, synexizw ston patera mou
            if level_state[level + 1] == 15:
                level_state[level + 1] = -1
                parent[level - 1]
                level -= 1
            # Alliws synexizw sto epomeno paidi
            else:
                level_state[level + 1] += 1
                parent[level + 1] = \
                    os.path.join(parent[level],
                                 hex_digits[level_state[level + 1]])
                level += 1


def next_subdir2():
    for d1 in hex_digits:
        for d2 in hex_digits:
            yield os.path.join(d1, d2)


def next_subdir3():
    for d1 in hex_digits:
        for d2 in hex_digits:
            for d3 in hex_digits:
                yield os.path.join(d1, d2, d3)


next_subdirX = next_subdir2 if INDEX_HEIGHT == 2 else next_subdir3


def main():
    for path in next_subdir():
        print path


if __name__ == '__main__':
    main()
