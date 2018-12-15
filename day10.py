from sys import stdin
from operator import itemgetter
import re

inp = [i.strip() for i in stdin.readlines()]

# Find the second when lights are closest together
# nums = [[int(i) for i in re.findall(r'-?\d+', l)] for l in inp]
# dif = list()
# for i in range(20000):
#     minx = min(x + i * vx for (x, y, vx, vy) in nums)
#     maxx = max(x + i * vx for (x, y, vx, vy) in nums)
#     miny = min(y + i * vy for (x, y, vx, vy) in nums)
#     maxy = max(y + i * vy for (x, y, vx, vy) in nums)
#     dif.append(maxx-minx+maxy-miny)
# print(dif.index(min(dif)))
# print(min(dif))

second = 10605
coords = [[tuple(map(int, j)) for j in re.compile('=<\s?(.*?),\s?(.*?)>').findall(i)] for i in inp]
pos, velo = list(), list()
for i in coords:
    pos.append(list(i[0]))
    velo.append(i[1])

for i in range(len(pos)):
        pos[i][0] += velo[i][0]*second
        pos[i][1] += velo[i][1]*second
maxi = max(max(pos, key=itemgetter(1)))
mini = min(min(pos, key=itemgetter(1)))
for i in range(mini-mini, maxi*2):
    for j in range(mini-mini, maxi*2):
        if [j,i] in pos:
            print('#', end='')
        else:
            print('.', end='')
    print()