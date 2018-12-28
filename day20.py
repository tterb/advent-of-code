from collections import defaultdict

dir = {
    "N": (0, -1),
    "E": (1, 0),
    "S": (0, 1),
    "W": (-1, 0)
}

inp = [i for i in input().strip()[1:-1]]
pos = list()
x, y = 500, 500
prev_x, prev_y = x, y
map = defaultdict(set)
distance = defaultdict(int)
for i in inp:
    if i == "(":
        pos.append((x, y))
    elif i == ")":
        x, y = pos.pop()
    elif i == "|":
        x, y = pos[-1]
    else:
        dx, dy = dir[i]
        x += dx
        y += dy
        map[(x, y)].add((prev_x, prev_y))
        if distance[(x, y)]:
            distance[(x, y)] = min(distance[(x, y)], distance[(prev_x, prev_y)]+1)
        else:
            distance[(x, y)] = distance[(prev_x, prev_y)]+1
    prev_x, prev_y = x, y
print(max(distance.values()))