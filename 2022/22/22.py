import fileinput
import re

input = list(fileinput.input())
m = {(x,y):c for y, line in enumerate(input) for x, c in enumerate(line) if c in '#.'}
path = re.findall('[LR]|[0-9]+', input[-1])
for i in range(0, len(path), 2): path[i] = int(path[i])
size = int(((len(m)/6)**0.5))
mid = (size-1)/2
dirs = [(1,0), (0,1), (-1,0), (0,-1)]
add = lambda dir, delta=0: (dir+delta+4)%4
offset = lambda pos, dir, steps=size: (pos[0]+steps*dirs[dir][0], pos[1]+steps*dirs[dir][1])
sides = {(x-x%size,y-y%size) for (x,y) in m}
glue2d = {(pos,d): (offset(pos,d),0) for pos in sides for d in range(4) if offset(pos,d) in sides}
glue3d = dict(glue2d)

# glue2d just wraps
for side, dir in ((side,dir) for side in sides for dir in range(4)):
    if (side,dir) not in glue2d:
        for back in range(0, 6*size, size):
            if offset(side, add(dir,2), back+size) not in m: break
        glue2d[(side,dir)] = (offset(side,add(dir,2),back), 0)

# If we can move forward and turn left twice, we should be facing ourself again.
def curve(pos, dir, recurse=True):
    side, delta = glue3d.get((pos, dir), (None,0))
    if not side or not recurse: return (side, add(dir, delta-1))
    return curve(side, add(dir, delta-1), False)

# We're done with glue3d when we didn't find any edges to glue together.
done = False
while not done:
    done = True
    for side, dir in ((side,dir) for side in sides for dir in range(4) if (side, dir) not in glue3d):
        done = False
        topSide, topDir = curve(side, add(dir,1))
        if topSide:
            delta = add(topDir+2, -dir)
            glue3d[(side, dir)] = (topSide, delta)
            glue3d[(topSide, topDir)] = (side, add(-delta))

def rotateTo(pos, side, delta):
    vecX, vecY = pos[0]%size - mid, pos[1]%size - mid
    newPos = (side[0]+mid, side[1]+mid)
    newPos = offset(newPos, delta, vecX)
    newPos = offset(newPos, add(delta,1), vecY)
    return (int(newPos[0]), int(newPos[1]))

def walk1(pos, dir, glue):
    proposed, newDir = offset(pos, dir, 1), dir
    if proposed not in m:
        newSide, delta = glue[((pos[0]//size*size,pos[1]//size*size), dir)]
        proposed = rotateTo(proposed, newSide, delta)
        newDir = add(dir, delta)

    if m[proposed] != '#':
        return proposed, newDir
    return pos, dir  

def walk(glue):
    pos = (min(p[0] for p in m if p[1]==0 and m[p] == '.' ), 0)
    dir = 0
    for cmd in path:
        if cmd == 'L':
            dir = add(dir,-1)
        elif cmd == 'R':
            dir = add(dir,1)
        else:
            for i in range(cmd):
                pos, dir = walk1(pos, dir, glue)
    print(1000*(pos[1]+1) + 4 * (pos[0]+1) + dir)

walk(glue2d)
walk(glue3d)

def draw(m):
    xmin, xmax = min(k[0] for k in m), max(k[0] for k in m)
    ymin, ymax = min(k[1] for k in m), max(k[1] for k in m)
    for y in range(ymin, ymax+1):
        for x in range(xmin,xmax+1):
            print(m.get((x,y), ' '), end='')
        print()
