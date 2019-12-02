from sys import stdin
import itertools

def getAdjacent(acres, pos):
  size = len(acres)
  for i in itertools.product(*(range(n-1, n+2) for n in pos)):
    if i != pos and all(0 <= n < size for n in i):
      yield i

def printAcres(acres, n):
  print(f'After {n} minutes:')
  for y in acres:
    for x in y:
      print(x, end='')
    print()
  print()
    
def main():
  minutes = 1000000000
  acres = [list(i.strip()) for i in stdin.readlines()]
  # printAcres(acres, n)
  size = len(acres)
  adj = dict()
  for y in range(size):
    for x in range(size):
      adj[(x,y)] = list(getAdjacent(acres, (x,y)))
      # print(f'{(x,y)} - {list(getAdjacent(acres, (x,y)))}')
  wooded = sum([i.count('|') for i in acres])
  lumber = sum([i.count('#') for i in acres])
  for n in range(1, minutes):
    next = [['.']*size for _ in range(size)]
    for y in range(size):
      for x in range(size):
        temp = [acres[i[1]][i[0]] for i in adj[(x,y)]]
        if acres[y][x] == '.' and temp.count('|') >= 3:
          next[y][x] = '|'
        elif acres[y][x] == '|' and temp.count('#') >= 3:
          next[y][x] = '#'
        elif acres[y][x] == '#' and (not temp.count('#') or not temp.count('|')):
          next[y][x] = '.'
        else:
          next[y][x] = acres[y][x]
    acres = next
  total = sum([i.count('|') for i in acres])*sum([i.count('#') for i in acres])
  print(total)
  
# Part 2:
# You're able to calculate the board state at 1000000000 minutes, after
# discovering that the board state repeats every 28 iterations after the first
# few hundred minutes.
  
if __name__ == '__main__':
  main()