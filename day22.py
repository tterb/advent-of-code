
def expand(g):
	for i in range(len(g)):
		g[i] = [0]+g[i]+[0];
	n = [0]*len(g[0]);
	return [n]+g+[n];

def move(n, pos, d):
	arr = ['up','right','down','left'];
	if n == 0:
		n -= 1;
	i = arr.index(d);
	d = arr[(i+n)%len(arr)]; 
	if d == 'up':
		pos = (pos[0], pos[1]+1);
	elif d == 'right': 
		pos = (pos[0]+1, pos[1]);
	elif d == 'down':
		pos = (pos[0], pos[1]-1);
	elif d == 'left':
		pos = (pos[0]-1, pos[1]);
	return pos, d;

def main():
	graph = [];
	while True:
		try: 
			line = input();
			n = [];
			for i in line:
				if i == '.':
					n.append(0);
				else:
					n.append(1)
			graph.append(n); 
		except EOFError:
			break;

	pos, direct = (int(len(graph)/2),int(len(graph)/2)), 'left';
	count, infected = 0, 0;
	while count < 10000:
		for x in pos:
			if x == len(graph[0])-1 or x == 0:
				graph = expand(graph);
				pos = (pos[0]+1, pos[1]+1)
		node = graph[pos[0]][pos[1]];
		infected += 1 - node; 
		graph[pos[0]][pos[1]] = 1 - node;
		pos, direct = move(node, pos, direct);
		count += 1;
	print(infected);


if __name__ == "__main__":
    main()