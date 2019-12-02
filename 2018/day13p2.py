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
count = 0
while len(carts) > 1:
    print(f'{count}: ')
    pos = [x[0] for x in carts]
    for n,i in enumerate(carts):
        print(f'{n}: ({i[0]}, {i[1]}, {i[2]}) =>', end=' ')
        i[0] = (i[0][0]+i[1][0], i[0][1]+i[1][1])
        next = grid[i[0][1]][i[0][0]]
        if i[0] in pos:
            cart.remove(i)
            for k in carts:
                if k[0] == temp:
                    carts.remove(k)
                    break
            # crash = i[0]
            # break
        else:
            if next == '+':
                i[1] = dir[(dir.index(i[1])+i[2]-1)%len(dir)]
                i[2] = (i[2]+1)%3
            elif (next == '/' and dir.index(i[1])%2) or (next == '\\' and not dir.index(i[1])%2):
                i[1] = dir[(dir.index(i[1])+1)%len(dir)]
            elif next == '/' or next == '\\':
                i[1] = dir[(dir.index(i[1])-1)%len(dir)]
            carts[n] = [i[0], i[1], i[2]]
            print(f'\'{next}\' => ({i[0]}, {i[1]}, {i[2]})')
    count += 1
print(crash)