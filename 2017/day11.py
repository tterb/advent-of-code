# The hexagons ("hexes") in this grid are aligned such that adjacent hexes can be found to the north, northeast, southeast, south, southwest, and northwest.
# You have a path, starting where he started, you need to determine the fewest number of steps required to reach him. (A "step" means to move from the hex you are in to any adjacent hex.)

val = {
	'n': (0,1,-1),
	's': (0,-1,1),
	'nw': (-1,1,0),
	'ne': (1,0,-1),
	'sw': (-1,0,1),
	'se': (1,-1,0)
}

x, y, z = 0, 0, 0;
directions = [x for x in input().split(',')];

for d in directions:
	x += val[d][0]; 
	y += val[d][1];
	z += val[d][2];
	
print(int((abs(x) + abs(y) + abs(z))/2));

