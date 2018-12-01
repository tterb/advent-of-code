
def expand(g):
	for i in range(len(g)):
		g[i] = [0]+g[i]+[0];
	n = [0]*len(g[0]);
	return [n]+g+[n];

def move(n, pos, d):
	arr = ['up','right','down','left'];
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

def printer(arr):
	for i in arr:
		[print(' '.join('.' if x == 0 else '#' for x in i))];

def partOne(graph, pos, direct):
	count, infected = 0, 0;
	while count < 70:
		for x in pos:
			if x == len(graph[0])-1 or x == 0:
				graph = expand(graph);
				pos = (pos[0]+1, pos[1]+1)
		node = graph[pos[0]][pos[1]];
		infected += 1 - node; 
		graph[pos[0]][pos[1]] = 1 - node;
		v = 1;
		if node == 0: 
			v = -1;
		pos, direct = move(v, pos, direct);
		count += 1;
	return infected;

def partTwo(graph, pos, direct):
	weak, flagged = set(), set();
	count, infected = 0, 0;
	while count < 10:
		for x in pos:
			if x == len(graph[0])-1 or x == 0:
				graph = expand(graph);
				pos = (pos[0]+1, pos[1]+1);
				weak = set([(t[0]+1,t[1]+1) for t in weak]);
				flagged = set([(t[0]+1,t[1]+1) for t in flagged]);
		node = graph[pos[0]][pos[1]];
		v = 0;
		if pos in weak:
			infected += 1 - node; 
			graph[pos[0]][pos[1]] = 1 - node;
			weak.remove(pos);
		elif node == 0: 
			weak.add(pos);
			v -= 1;
		elif pos in flagged:
			infected += 1 - node; 
			graph[pos[0]][pos[1]] = 1 - node;
			flagged.remove(pos);
			v += 2;
		elif node == 1:
			flagged.add(pos);
			v += 1;
		pos, direct = move(v, pos, direct);
		count += 1;
	return infected;



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
	pos = (int(len(graph)/2),int(len(graph)/2));
	# print(partOne(graph, pos, direct));
	print(partTwo(graph, pos, 'left'));
	exit();


if __name__ == "__main__":
    main()