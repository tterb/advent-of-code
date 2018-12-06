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
max = 0
occurances = dict()
for i in guards.keys():
  counter = Counter(guards[i])
  occurances[i] = counter.most_common(1)[0]
  if occurances[i][1] > max:
    maxId, minute, max = i, occurances[i][0], occurances[i][1]
print(maxId*minute)