from sys import stdin
from collections import defaultdict

coords = defaultdict(int)
for claim in stdin.readlines():
    claim = claim.split(' @ ')[-1].strip()
    x, y = [int(i) for i in claim.split(': ')[0].split(',')]
    w, h = [int(i) for i in claim.split(': ')[1].split('x')]
    for i in range(w):
        for j in range(h):
            coords[(x+i,y+j)] += 1
print(sum(1 for i in coords.values() if i > 1))