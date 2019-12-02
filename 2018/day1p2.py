import sys

inp = [int(i) for i in sys.stdin.readlines()]
total = 0
vals = set()
while True:
    for i in inp:
        vals.add(total)
        total += i
        if total in vals:
            print(total)
            sys.exit()
        vals.add(total)