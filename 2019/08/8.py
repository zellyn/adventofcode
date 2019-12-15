#! /usr/bin/env python3

import functools

def split(s, l): return [s[i:i+l] for i in range(0, len(s), l)]

def merge(f, b): return ''.join([x[x[0]=='2'] for x in zip(f, b)])

with open('input', 'r') as f:
    input = f.read().strip()
    parts = split(input, 25*6)
    lowest = min(parts, key=lambda x:x.count('0'))
    print(lowest.count('1') * lowest.count('2'))

    image = functools.reduce(merge, parts)

    for line in split(image, 25):
        print(line.replace("1", "\033[7m1\033[m"))
