import numpy as np

n = 300
serial = int(input())
grid = np.array([[int(str(((x+10)*y+serial)*(x+10))[-3])-5 for y in range(1, n+1)] for x in range(1, n+1)])
coord = (0, 0)
mVal, dim = 0, 0
for d in range(4, 2, -1):
    squares = sum(grid[x:x-d+1 or None, y:y-d+1 or None] for x in range(d) for y in range(d))
    val = int(squares.max())
    if mVal < val:
        coord = np.where(squares == val)
        mVal = val
        dim = d
x,y = coord[0][0], coord[1][0]
print(f'({x+1}, {y+1}) X {dim} = {mVal}')