from sys import stdin

class Disjoint:
    def __init__(self, id):
        self.id = id
        self.parent = self
        self.rank = 0
    def __str__(self):
        return '('+','.join(map(str, self.id))+')'

def find(x):
    if x.parent is not x:
        x.parent = find(x.parent)
    return x.parent

def union(x, y):
    xroot = find(x)
    yroot = find(y)
    # x and y are already in same set
    if xroot == yroot:
        return
    # x and y are not in same set, so we merge them
    if xroot.rank < yroot.rank:
        xroot, yroot = yroot, xroot
    yroot.parent = xroot
    if xroot.rank == yroot.rank:
        xroot.rank += 1

def distance(a,b):
    return abs(a[0]-b[0])+abs(a[1]-b[1])+abs(a[2]-b[2])+abs(a[3]-b[3])

points = [Disjoint(tuple(int(j) for j in i.strip().split(','))) for i in stdin.readlines()]
for i in points:
    for j in points:
        if distance(i.id, j.id) <= 3:
            union(i, j)
print(sum([1 for i in points if i.parent == i]))
