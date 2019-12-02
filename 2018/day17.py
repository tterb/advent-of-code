from sys import stdin
from itertools import groupby
from math import inf
import re

def printClay(clay):
    minX = min(clay, key=lambda x: x[0])[0]
    maxX = max(clay, key=lambda x: x[0])[0]
    maxY = max(clay, key=lambda x: x[1])[1]
    center = (maxX - minX)//2
    for y in range(maxY+1):
        for x in range(minX, maxX+1):
            if (x,y) == (center+minX, 0):
                print('+', end='')
            elif (x,y) in clay:
                print('#', end='')
            else:
                print('.', end='')
        print()
    print()
    
def printGrid(grid):
    print()
    for y in range(len(grid)):
        for x in range(len(grid[0])):
            if (x,y) == (len(grid[0])//2, 0):
                print('+', end='')
            else:
                print(grid[y][x], end='')
        print()
    print()

def fillBlocks(grid, center, bottom, total):
    groups = [(k, sum(1 for i in g)) for k,g in groupby(grid[bottom])]
    index, width = 0, 0
    row = bottom-1
    for i in groups:
        if index+i[1] >= center:
            width = i[1]
            break
        index += i[1]
    while grid[row][index+width-1] == '#':
        leftX, rightX = center, center
        while(leftX > index or rightX < index+width) and (grid[row][leftX] != '#' or grid[row][rightX] != '#'):
            if leftX > index and grid[row][leftX] != '#':
                grid[row][leftX] = '~'
                leftX -= 1
            if rightX < index+width and grid[row][rightX] != '#':
                grid[row][rightX] = '~'
                rightX += 1
        print(f'({leftX}, {rightX})')
        total += (rightX-leftX)-1
        # total += row
        row -= 1
    print(row)
    # Overflow from resevoir
    # if grid[row-1][leftX] == '#':
    #     leftX += 1
    # if grid[row-1][rightX] == '#':
    #     rightX -= 1
    # total += (rightX-leftX)-1
    return total, rightX

clay = list()
for line in stdin.readlines():
    mX = inf 
    a,b,c = map(int, re.findall("([\d]+)", line.strip()))
    x, y = list(), list()
    if line[0] == 'x':
        for i in range(b, c+1):
            clay.append((a, i))
    else:
        for i in range(b, c+1):
            clay.append((i, a))
printClay(clay)
minX = min(clay, key=lambda x: x[0])[0]
maxX = max(clay, key=lambda x: x[0])[0]
maxY = max(clay, key=lambda x: x[1])[1]
center = (maxX - minX)//2
grid = [['.' for j in range((maxX-minX)+1)] for i in range(maxY+1)]
for y in range(maxY+1):
    for x in range((maxX-minX)+1):
        if (x+minX,y) in clay:
            grid[y][x] = '#'
bottom = [grid[n][center] for n in range(maxY+1)].index('#')
print(bottom)
total, center = fillBlocks(grid, center, bottom, 0)
printGrid(grid)
print(f'center: {center}')
print(f'total: {total}')
printGrid(grid)
tempGrid = grid[bottom+1:]
temp = bottom+1
bottom = [tempGrid[n][center] for n in range(maxY+1-temp)].index('#')
total, center = fillBlocks(tempGrid, center, bottom, total)
print(f'center: {center}')
print(f'total: {total}')

printGrid(grid)
print(f'center: {center}')
# print(f'width: {width}')
# print(f'height: {height}')
print(f'total: {total}')


    