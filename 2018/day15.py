from sys import stdin
from collections import defaultdict
from math import inf
import itertools

class Actor:
    def __init__(self, type, pos):
        self.isElf = type
        self.pos = pos
        self.hp = 200
        self.dead = False
    # def __eq__(self, a):
    #     return (self.pos == a.pos) and (self.type == type) 
    def __str__(self):
        return '('+str(self.pos[0])+', '+str(self.pos[1])+')'
    
def updateDist(dist):
    for i in dist:
        for j in dist[i]:
            j = distance(k[0].pos, k[1].pos)
    return dist

def distance(a, b):
    return abs(a[0]-b[0])+abs(a[1]-b[1])

def getMove(grid, dist, a):
    directions = [(-1,0),(0,-1),(0,1),(1,0)]
    for b in sorted(dist[a], key=lambda k: distance(a.pos, k.pos)):
        if a.isElf == b.isElf or a.dead or b.dead:
            continue
        # print(f'{a.pos} : {b.pos} : {dist[a][b]}')
        opts = [((a.pos[0]+j[0],a.pos[1]+j[1]), distance((a.pos[0]+j[0],a.pos[1]+j[1]), b.pos)) for j in directions]
        target = 'G'
        if not a.isElf:
            target = 'E'
        # if all(grid[i[0][0]][i[0][1]] == target for i in opts):
        #    print('ATTACK!!')
           # exit()
        print(f'opts: {sorted(opts, key=lambda j: j[1])}')
        vals = [k[1] for k in opts]
        if 0 in vals:
            for d in sorted(opts, key=lambda j: j[1]):
                if d[1] > 0:
                    break
                b.hp -= 3
                if b.hp <= 0:
                    b.dead = True
                    grid[b.pos[0]][b.pos[1]] = '.'
                break
        for d in sorted(opts, key=lambda j: j[1]):
            # print(f'{d}, ', end='')
            if grid[d[0][0]][d[0][1]] == '.': # and d[1] < dist[a][b]:
                print(f'{a.pos} => {d[0]}')
                # Update grid
                grid[a.pos[0]][a.pos[1]] = '.'
                if a.isElf:
                    grid[d[0][0]][d[0][1]] = 'E'
                else:
                    grid[d[0][0]][d[0][1]] = 'G'
                a.pos = d[0]
                dist[a][b] = d[1]
                dist[b][a] = d[1]
                return
        print()

def printGrid(grid, actors):
    for y,i in enumerate(grid):
        coords = list()
        for x,j in enumerate(i):
            if j == 'G' or j == 'E':
                coords.append((y,x))
            print(j, end='')
        if len(coords):
            points = list()
            for b in coords:
                for a in actors:
                    if a.pos == b:
                        points.append(a.hp)
            print(f'  {points}')
        else:
            print()
    print()
    
def main():
    grid = [list(i.strip()) for i in stdin.readlines()]
    actors = list()
    dist = defaultdict(lambda: defaultdict(lambda: inf))
    for y in range(len(grid)):
        for x in range(len(grid[y])):
            if grid[y][x] == 'E' or grid[y][x] == 'G':
                actors.append(Actor((grid[y][x] == 'E'), [y,x]))
    for a,b in itertools.combinations(actors, 2):
        d = distance(a.pos, b.pos)
        dist[a][b] = d
        dist[b][a] = d
    printGrid(grid, actors)
    for i in range(4):
        for a in actors:
            getMove(grid, dist, a)
        actors.sort(key=lambda x: x.pos[0])
        [print(k, end=', ') for k in actors]
        print()
        printGrid(grid, actors)
        # dist = updateDist(dist)

if __name__ == '__main__':
    main()