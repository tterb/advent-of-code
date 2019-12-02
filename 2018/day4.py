from sys import stdin
from collections import defaultdict
from collections import Counter
import re

inp = [i.strip() for i in stdin.readlines()]
inp.sort()
guards = defaultdict(list)
i = 0
while(i < len(inp)):
  if re.search('#(\d+)', inp[i]):
    num = int(re.search('#(\d+)', inp[i]).group()[1:])
    i += 1
    while(i < len(inp) and not re.search('#(\d+)', inp[i])):
      guards[num] += list(range(int(inp[i].split(':')[1][:2]), int(inp[i+1].split(':')[1][:2])))
      i += 2
# [print(f'{i}: {guards[i]}') for i in guards.keys()]
id = max(guards, key=lambda x: len(guards[x]))
counter = Counter(guards[id])
common = counter.most_common(1)[0][0]
# print(f'{id} * {common}')
print(id*common)