import networkx as nx

def printCave(cave, target, dim):
  for y in range(dim[1]):
    for x in range(dim[0]):
      if (x,y) == target:
        print('T', end='')
      elif (x,y) == (0,0):
        print('M', end='')
      elif cave[y][x] == 0:
        print('.', end='')
      elif cave[y][x] == 1:
        print('=', end='')
      elif cave[y][x] == 2:
        print('|', end='')
    print()
  print()
  
def generatecave(dim):
  cave = dict()
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
      cave[(x,y)] = ero[(x,y)]%3
  return cave

def dijkstra(cave, dim, tools, target):
    graph = nx.Graph()
    for y in range(dim[1]+1):
      for x in range(dim[0]+1):
        tool = tools[cave[(x, y)]]
        graph.add_edge((x, y, tool[0]), (x, y, tool[1]), weight=7)
        for dx, dy in ((0, 1), (0, -1), (1, 0), (-1, 0)):
          nextX, nextY = x+dx, y+dy
          if 0 <= nextX <= dim[0] and 0 <= nextY <= dim[1]:
            opts = tools[cave[(nextX, nextY)]]
            for i in set(tools).intersection(set(opts)):
                graph.add_edge((x, y, i), (nextX, nextY, i), weight=1)
    return nx.dijkstra_path_length(graph, (0, 0, 0), (target[0], target[1], 0))

def calcRisk(cave, target):
  total = 0
  for y in range(target[1]+1):
    total += sum(cave[(x,y)] for x in range(target[0]+1))
  return total


depth = int(input().split(': ')[-1])
target = tuple(map(int,input().split(': ')[-1].split(',')))
dim = (target[0]+100, target[1]+100)
tools = [('climbing','torch'),('climbing','neither'),('torch','neither')]
cave = generatecave(dim)
# printCave(cave, target, dim)
print(f'Part 1: {calcRisk(cave, target)}')

cave = {i:j for i,j in generatecave(dim).items()}
print(f'Part 2: {dijkstra(cave, dim, tools, target)}')
