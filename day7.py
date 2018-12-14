from sys import stdin

class Node:
	def __init__(self, data):
		self.data = data
		self.prereqs = set()
		self.adjacent = list()

def topoDFS(graph, visited, opts, path, node):
	path.append(node)
	visited.add(node)
	opts = sorted(list(set(node.adjacent+opts)), key=lambda x: x.data)
	while len(opts):
		n = opts.pop(0)
		if n not in visited and n.prereqs.issubset(visited):
			topoDFS(graph, visited, opts, path, n)
	return path

def main():
	steps = [(j[1],j[7]) for j in [i.strip().split(' ') for i in stdin.readlines()]]
	nodes = dict()
	for i in steps:
		if i[0] not in nodes:
			nodes[i[0]] = Node(i[0])
		if i[1] not in nodes:
			nodes[i[1]] = Node(i[1])
		nodes[i[0]].adjacent.append(nodes[i[1]])
		nodes[i[1]].prereqs.add(nodes[i[0]])
	opts = list(set(nodes[i[0]] for i in steps).difference(set(nodes[i[1]] for i in steps)))
	opts.sort(key=lambda x: x.data)
	[print(n.data, end=',') for n in opts]
	print()
	visited = set()
	path = topoDFS(nodes, set(), opts, list(), opts.pop(0))
	[print(n.data, end='') for n in path]


if __name__ == '__main__':
    main()