# The disk in question consists of a 128x128 grid; each square of the grid is either free or used. On this disk, the state of the grid is tracked by the bits in a sequence of knot hashes.
# A total of 128 knot hashes are calculated, each corresponding to a single row in the grid; each hash contains 128 bits which correspond to individual grid squares. Each bit of a hash indicates whether that square is free (0) or used (1).


import day10 as ten;
RNG = 128;

inp = input();
knotOut = [ten.partTwo(inp+'-'+str(i)) for i in range(RNG)];
binOut = [bin(int(i, 16))[2:].zfill(4) for i in knotOut];
count = 0;
for i in binOut:
	count += i.count('1');
print(count);