
def printMap(map, target, dim):
    for y in range(dim[1]):
        for x in range(dim[0]):
            if (x,y) == target:
                print('T', end='')
            elif (x,y) == (0,0):
                print('M', end='')
            elif map[y][x] == 0:
                print('.', end='')
            elif map[y][x] == 1:
                print('=', end='')
            elif map[y][x] == 2:
                print('|', end='')
        print()
    print()

def calcRisk(map, target):
    total = 0
    for y in range(target[1]+1):
        total += sum(map[y][x] for x in range(target[0]+1))
    return total


depth = int(input().split(': ')[-1])
target = tuple(map(int,input().split(': ')[-1].split(',')))
dim = (int(target[0]*(1.5))+1, int(target[1]*(1.5))+1)
map = [[0 for i in range(dim[0])] for _ in range(dim[1])]
ero = dict()
for y in range(dim[1]):
    for x in range(dim[0]):
        if (x,y) in [(0,0), target]:
            geo = 0
        elif y == 0:
            geo = x*16807
        elif x == 0:
            geo = y*48271
        else:
            geo = ero[(x-1,y)]*ero[(x,y-1)]
        ero[(x,y)] = (geo+depth)%20183
        map[y][x] = (ero[(x,y)])%3
# printMap(map, target, dim)
risk = calcRisk(map, target)
print(risk)