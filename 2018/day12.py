from sys import stdin
from collections import defaultdict

n = 102
state = ([0]*5)+[int(i == '#') for i in stdin.readline().split(' ')[-1]]+([0]*n*2)
input()
pattern = defaultdict(lambda: False)
for i in stdin.readlines():
    p = i.strip().split(' => ')
    gen = [int(j =='#') for j in p[0]]
    pattern[''.join(map(str, gen))] = (p[1] == '#')
# print(f'0: {"".join(["#" if i else "." for i in state])}')
for i in range(1, n+1):
    temp = []+state
    for j in range(2, len(state)):
        temp[j] = int(pattern[''.join(map(str, state[j-2:j+3]))])
    state = temp
    # print(f'{i}: {"".join(["#" if i else "." for i in state])})

# part 1
# print(sum(i-5 for i in range(len(state)) if state[i]))

# part 2
total = sum(i-5 for i in range(len(state)) if state[i])
print(((50000000000-102)*46)+total)