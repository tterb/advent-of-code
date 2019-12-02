from sys import stdin
            
grid = [list(i.rstrip()) for i in stdin.readlines()]
carts = list()
arrows = ['<','^','>','v']
dir = [(-1,0),(0,-1),(1,0),(0,1)]
for y in range(len(grid)):
    for x,i in enumerate(grid[y]):
        if i in arrows:
            carts.append([(x,y), dir[arrows.index(i)], 0])
            if i in ['^','v']:
                grid[y][x] = '|'
            else:
                grid[y][x] = '-'
crash = None
while not crash:
    pos = [x[0] for x in carts]
    for n,i in enumerate(carts):
        i[0] = (i[0][0]+i[1][0], i[0][1]+i[1][1])
        next = grid[i[0][1]][i[0][0]]
        if i[0] in pos:
            crash = i[0]
            break
        if next == '+':
            i[1] = dir[(dir.index(i[1])+i[2]-1)%len(dir)]
            i[2] = (i[2]+1)%3
        elif (next == '/' and dir.index(i[1])%2) or (next == '\\' and not dir.index(i[1])%2):
            i[1] = dir[(dir.index(i[1])+1)%len(dir)]
        elif next == '/' or next == '\\':
            i[1] = dir[(dir.index(i[1])-1)%len(dir)]
        carts[n] = [i[0], i[1], i[2]]
print(crash)