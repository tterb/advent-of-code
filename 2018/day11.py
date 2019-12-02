
n = 300
serial = int(input())
grid = [[int(str(((x+10)*y+serial)*(x+10))[-3])-5 for y in range(1, n+1)] for x in range(1, n+1)]
coord, val = (0, 0), 0
for x in range(n-2):
    for y in range(n-2):
        power = sum([grid[x+i][y+j] for j in range(3) for i in range(3)])
        if power > val:
            coord, val = (x+1, y+1), power
print(f'{coord} = {val}')