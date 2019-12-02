from sys import stdin
from queue import PriorityQueue
import re

# Manhattan distance
def distance(a,b):
  return abs(a[0]-b[0])+abs(a[1]-b[1])+abs(a[2]-b[2])

nanos = list()
for i in stdin.readlines():
  pattern = re.compile('<(.*\d),(.*\d),(.*\d)>, r=(.*\d)')
  x,y,z,r = map(int, pattern.findall(i.strip())[0])
  nanos.append([(x,y,z), r])
strong = max(nanos, key=lambda x: x[1])
count = sum([distance(strong[0], n[0]) <= strong[1] for n in nanos])
print(f'Part 1: {count}')

q = PriorityQueue()
for pos,r in nanos:
  dist = abs(pos[0])+abs(pos[1])+abs(pos[2])
  q.put((max(0, dist-r), 1))
  q.put((dist+r+1, -1))
count, maxCount, result = 0, 0, 0
while not q.empty():
  dist, n = q.get()
  count += n
  if count > maxCount:
    maxCount, result = count, dist
print(f'Part 2: {result}')