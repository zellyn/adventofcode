import re
import sys

filename = sys.argv[1]
row = int(sys.argv[2])
limit = int(sys.argv[3])

def dist(x1, y1, x2, y2):
    return abs(x2-x1) + abs(y2-y1)

def overlaps(x,y, sensor):
    sx, sy, r, minx, maxx, miny, maxy = sensor
    return minx<=x and x<=maxx and miny<=y and y<=maxy and dist(x, y, sx, sy) <= r

sensors = []
beacons_in_row = set()
for line in open(filename).readlines():
    x, y, bx, by = [int(x) for x in re.split(r"[^0-9-]+", line) if x]
    r = dist(x,y,bx,by)
    sensors.append((x, y, r, x-r, x+r, y-r, y+r))
    if by == row:
        beacons_in_row.add(bx)

minx = min(s[0]-s[2] for s in sensors)
maxx = max(s[0]+s[2] for s in sensors)
print(sum([any(overlaps(x,row,s) for s in sensors) for x in range(minx,maxx+1)]) - len(beacons_in_row))

count = [0]
def edges(x,y,r, limit):
    count[0] += 1
    print("Sensor", count[0])
    top, bottom = y-r, y+r

    for i in range(max(0,-top), min(r,limit-x-1)+1):
        yield (x+1+i, top+i)
    for i in range(max(0,-top), min(r, x-1)+1):
        yield (x-1-i, top+i)
    for i in range(max(0,bottom-limit), min(r,limit-x)):
        yield (x+1+i, bottom-i)
    for i in range(max(0,bottom-limit), min(r, x)):
        yield (x-1-i, bottom-i)

positions = ((x,y) for i in (edges(s[0], s[1], s[2], limit) for s in sensors) for (x,y) in i)
found = [(x,y) for (x,y) in positions if all(not overlaps(x,y,s) for s in sensors)]
(x,y) = found[0]
print(x*4000000+y)
