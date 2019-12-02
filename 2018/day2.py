from sys import stdin
from collections import Counter

two, three = 0, 0
for id in stdin.readlines():
    counter = Counter(id)
    if 2 in counter.values():
        two += 1
    if 3 in counter.values():
        three += 1
print(two*three)
    
    