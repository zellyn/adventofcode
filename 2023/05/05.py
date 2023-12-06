#!/usr/bin/env python3

import re

with open('input') as f:
    input = [[[int(x) for x in line.split()]
              for line in piece.strip().split('\n')]
             for piece in re.split(r'(?m)\n*.*: *\n*', f.read())[1:]]
    seeds = input[0][0]
    maps = [sorted([(s,s+l-1,d-s) for d,s,l in mp]) for mp in input[1:]]

def split_at_all(range, map_ranges):
    result = []
    for split_point in sorted(set([x[0] for x in map_ranges] + [x[1]+1 for x in map_ranges])):
        if split_point <= range[0]: continue
        if split_point > range[1]:
            return result + [range]
        result.append((range[0], split_point-1))
        range = (split_point, range[1])
    return result + [range]

def map_single_range(range, map_ranges):
    ranges = sorted(split_at_all(range, map_ranges))
    map_ranges = map_ranges + [(float('inf'), float('inf'), 0)]
    result = []
    for range in ranges:
        while map_ranges[0][1] < range[1]: map_ranges.pop(0)
        if range[1] < map_ranges[0][0]:
            result.append(range)
        else:
            result.append((range[0]+map_ranges[0][2], range[1]+map_ranges[0][2]))
    return result

def lowest(ranges, maps):
    for map_ranges in maps:
        new_ranges = []
        for range in ranges:
            new_ranges.extend(map_single_range(range, map_ranges))
        ranges = new_ranges
    return sorted(ranges)[0][0]

print(lowest(list(zip(seeds, seeds)), maps))
print(lowest([(x,x+y-1) for x,y in zip(seeds[0::2], seeds[1::2])], maps))
