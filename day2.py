from sys import stdin
from collections import Counter

validCount = 0
for line in stdin.readlines():
    rng, val, pwd = line.split(' ')
    minCount, maxCount = map(int, rng.split('-'))
    char = val.replace(':', '')
    if minCount <= pwd.count(char) <= maxCount:
        validCount += 1
print(f'Valid passwords: {validCount}')
