# Programs in this village communicate using a fixed system of pipes. Messages are passed between programs using these pipes, but most programs aren't connected to each other directly. Instead, programs pass messages between each other until the message reaches the intended recipient.
# You need to figure out how many programs are in the group that contains program ID 0.

import re;

def dfs(pipes, visited, u):
	for k in pipes[u]:
		if k not in visited:
			visited.add(k);
			dfs(pipes, visited, k);
	return visited;


def main():
	pipes = {};
	visited = set();
	while True:
		try:
			inp = re.findall(r"[\w']+", input()); 
			inp = set([int(s) for s in inp]);
			for i in inp:
				if i in pipes.keys():
					pipes[i] = inp.union(set(pipes[i]));
				else:
					pipes[i] = set(inp)
		except EOFError:
			break;
	print(len(dfs(pipes, visited, 0)));


if __name__ == "__main__":
    main()
