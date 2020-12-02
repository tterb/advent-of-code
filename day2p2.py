from sys import stdin
from collections import Counter

validCount = 0
for line in stdin.readlines():
    rng, target, pwd = line.split(' ')
    minIndex, maxIndex = map(int, rng.split('-'))
    char = target.replace(':', '')
    if (pwd[minIndex-1] == char) ^ (pwd[maxIndex-1] == char):
        validCount += 1
print(f'Valid passwords: {validCount}')
