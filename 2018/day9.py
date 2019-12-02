from collections import deque

inp = input().split(' ')
players = int(inp[0])
# part 1
last = int(inp[6])
# part 2
# last = int(inp[6])*100
scores = [0 for i in range(players)]
marbles = deque([0])
for i in range(1, last+1):
    if i%23 == 0:
        marbles.rotate(7)
        scores[i%players] += i+marbles.pop()
        marbles.rotate(-1)
    else:
        marbles.rotate(-1)
        marbles.append(i)
print(max(scores))