from sys import stdin

vals = {}
for i in stdin.readlines():
    num = int(i)
    target = 2020 - num
    vals[num] = target
    if target in vals:
        print(target * vals[target])
        exit(0)
    vals[num] = target