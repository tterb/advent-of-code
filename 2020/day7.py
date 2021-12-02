from sys import stdin
from collections import defaultdict
from pprint import pprint


def get_containers(vertices, edges, target):
    containers = set([])
    for src in vertices:
        if src == target:
            continue
        visited = {key:False for key in vertices}
        # dfs(src, target, edges, visited, [])
        [containers.add(i) for i in dfs(src, target, edges, visited, [])]
    if target in containers:
        containers.remove(target)
    return containers

    
def dfs(src, dest, edges, visited, path):
    visited[src] = True
    path.append(src)
    if src == dest:
        return path
    elif len(edges[src]):
        for i in edges[src]:
            if not visited[i]:
                return dfs(i, dest, edges, visited, path)
    else:
        return []
    path.pop()
    visited[src] = False
    return []

def print_bags(edges):
    for k in edges.keys():
       print(f'{k}: {edges[k]}') 

def main():
    target = 'shiny gold'
    vertices = set()
    edges = defaultdict(list)
    for line in stdin.readlines():
        words = line.split(' ')
        vertex = f'{words[0]} {words[1]}'
        vertices.add(vertex)
        index = 4
        while (index + 3) < len(words):
            color = f'{words[index+1]} {words[index+2]}'
            edges[vertex].append(color)
            vertices.add(color)
            index += 4
    print_bags(edges)
    containers = get_containers(vertices, edges, target)
    for i in list(containers):
        containers.union(get_containers(vertices, edges, i))
    pprint(containers)
    print(f'Number of containers: {len(containers)}')



if __name__ == "__main__":
    main()